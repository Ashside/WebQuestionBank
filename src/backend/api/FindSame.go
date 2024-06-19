package api

import (
	"strings"
	"gorm.io/gorm"
	"log"
)
func FindSamePost(c *gin.Context) {
	var form struct {
		Username string `form:"username" binding:"required"`
		Id 		 int `form:"id" binding:"required"`
	}
	//获取id列表
	inputIDs :=queryTestByID(Id)

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "reason": "Invalid input format"})
		return
	}

	db, err := getDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Failed to connect to database"})
		return
	}

	// 查询用户是否存在
	var user Users
	if err := GetUserByUsername(db, form.Username, &user); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users not found"})
		return
	}

	// 鉴权，学生无法操作
	if user.Type == STUDENT {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Permission denied"})
		return

	}

	//分割
	var choiceIDs []int
	var subjectiveIDs []int

	for _, id := range inputIDs {
		questionType, err := getQuestionTypeByID(db,id)
		switch questionType {
		case "choice":
			choiceIDs = append(choiceIDs, id)
		case "subjective":
			subjectiveIDs = append(subjectiveIDs, id)
		default:
			log.Printf("Unknown question type for ID %d", id)
		}
	}

	// 查询不重复的选择题ID
	var distinctChoiceIDs []int
	if err := db.Table("choice_questions").Select("id").Where("id NOT IN ?", choiceIDs).Scan(&distinctChoiceIDs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Failed to fetch distinct choice question IDs"})
		return
	}

	// 查询不重复的主观题ID
	var distinctSubjectiveIDs []int
	if err := db.Table("subjective_questions").Select("id").Where("id NOT IN ?", subjectiveIDs).Scan(&distinctSubjectiveIDs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Failed to fetch distinct subjective question IDs"})
		return
	}

	ChoicesNodes :=MaxFlow(db *gorm.DB, choiceIDs []int, distinctChoiceIDs []int)
	SubjectiveNodes :=MaxFlow(db *gorm.DB, subjectiveIDs []int, distinctSubjectiveIDs []int)
	
	Choices_questions :=getQuestionsByTypeID(db, "Choice", ChoicesNodes)
	Subjective_questions :=getQuestionsByTypeID(db, "Subjective", SubjectiveNodess)
	
	questions := append(choiceQuestions, subjectiveQuestions...)
	c.JSON(http.StatusOK, gin.H{"success": true, "questions": questions})
}

func getQuestionTypeByID(db *gorm.DB, id int) (string, error) {
	// 尝试查询选择题
	var choice QuestionSummary
	if err := db.Table("choice_questions").Where("id = ?", id).Take(&choice).Error; err == nil {
		return "choice", nil
	} else if !gorm.IsRecordNotFoundError(err) {
		return "", err 
	}

	// 如果不是选择题，则尝试查询主观题
	var subjective QuestionSummary
	if err := db.Table("subjective_questions").Where("id = ?", id).Take(&subjective).Error; err == nil {
		return "subjective", nil
	} else if !gorm.IsRecordNotFoundError(err) {
		return "", err 
	}

	// 未知类型
	return "unknown", nil
}

func GetAllKeywords(db *gorm.DB, questionIDs []int) map[int][]string {
	questionKeywordsMap := make(map[int][]string)
	for _, qID := range questionIDs {
		var keywords []string
		db.Table("choice_questions").Joins("JOIN choice_question_keywords ON choice_questions.id = choice_question_keywords.question_id").
		Joins("JOIN keywords ON choice_question_keywords.keyword_id = keywords.id").
		Where("choice_questions.id = ?", qID).
		Pluck("keywords.keyword", &keywords)
		questionKeywordsMap[qID] = keywords
	}
	return questionKeywordsMap
}

func calculateSharedKeywordCount(setA, setB []string) int {
	setAMap := make(map[string]bool)
	for _, keyword := range setA {
		setAMap[keyword] = true
	}

	count := 0
	for _, keyword := range setB {
		if setAMap[keyword] {
			count++
		}
	}
	return count
}

type Edge struct {
	V   int // 目标顶点
	Cap int // 剩余容量
	Cost int // 单位流量通过费用
	Rev  int // 反向边索引
}

type Graph struct {
	List [][]Edge // 邻接表存储图
	N    int      // 顶点数
}
func NewGraph(n int) *Graph {
	return &Graph{
		List: make([][]Edge, n+10),
		N:    n,
	}
}
// addEdge 添加有向边u->v，容量cap，费用cost
func (g *Graph) addEdge(u, v, cap, cost int) {
	g.List[u] = append(g.List[u], Edge{V: v, Cap: cap, Cost: cost, Rev: len(g.List[v])})
	g.List[v] = append(g.List[v], Edge{V: u, Cap: 0, Cost: -cost, Rev: len(g.List[u])-1}) // 反向边
}

func MaxFlow(db *gorm.DB, IDs []int, distinctIDs []int) []int {
	var num int
	for _, fromID := range IDs {
		num=max(num,fromID)
	}
	for _, toID := range distinctIDs {
		num=max(num,toID)
	}

	g := NewGraph(num + 2)
	s := num+1 // 源点
	t := num+2 // 汇点

	for _, fromID := range IDs {
		g.addEdge(s, fromID, 1, 0)
	}
	for _, toID := range distinctIDs {
		g.addEdge(toID, t, 1, 0)
	}

	questionKeywordsMap := GetAllKeywords(db, append(IDs, distinctIDs...))

	for _, fromID := range IDs {
		for _, toID := range distinctIDs {
			go func(fid, tid int) {
				defer wg.Done()
				similarity := calculateSharedKeywordCount(questionKeywordsMap[fid], questionKeywordsMap[tid])
				if similarity > 0 {
					g.addEdge(fid,tid,1,similarity+5)
				} else{
					g.addEdge(fid,tid,1,1)
				}
			}(fromID, toID)
		}
	}

	flow, cost := g.dinic(s, t)
	// fmt.Printf("最大流: %d, 最小费用: %d\n", flow, cost)
	rightnodes :=RightNodes(Graph, t)
	return rightnodes
}

func (g *Graph) dinicDfs(v, t, f int, level []int, iter []int, currentCost int) (int, int) {
	if v == t {
		return f, currentCost * f // 累加当前路径的费用
	}
	for i := iter[v]; i < len(g.List[v]); i++ {
		e := &g.List[v][i]
		if e.Cap > 0 && level[v] < level[e.V] {
			d, pathCost := g.dinicDfs(e.V, t, min(f, e.Cap), level, iter, currentCost+e.Cost)
			if d > 0 {
				e.Cap -= d
				g.List[e.V][e.Rev].Cap += d
				return d, pathCost // 返回当前路径的流量和费用
			}
		}
	}
	return 0, 0
}

func (g *Graph) dinic(s, t int) (int, int) {
	flow, cost := 0, 0
	for {
		level := make([]int, g.N+10)
		iter := make([]int, g.N+10)
		maxLevel := g.bfs(s, t, level)
		if maxLevel == -1 {
			break
		}
		for {
			f, c := g.dinicDfs(s, t, 1<<60, level, iter, 0)
			if f == 0 {
				break
			}
			flow += f
			cost += c 
		}
	}
	return flow, cost
}

func (g *Graph) bfs(s, t int, level []int) int {
	for i := range level {
		level[i] = -1
	}
	level[s] = 0
	queue := []int{s}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, e := range g.List[v] {
			if e.Cap > 0 && level[e.V] < 0 {
				level[e.V] = level[v] + 1
				queue = append(queue, e.V)
			}
		}
	}
	return level[t]
}
func (g *Graph)RightNodes(t int) []int {
	var rightnodes []int

	for u := 0; u < g.N; u++ {
		if u == t {
			continue 
		}

		// 检查该节点到汇点t的边是否存在于邻接表中，且剩余容量为0
		found := false
		for _, edge := range g.List[u] {
			if edge.V == t && edge.Cap == 0 {
				found = true
				break
			}
		}

		// 如果找到了这样的边，将节点u加入结果列表
		if found {
			rightnodes = append(rightnodes, u)
		}
	}

	return rightnodes
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
type QuestionSummary struct {
	ID            int    `json:"id"`
	QuestionType  string `json:"question_type"`
	Subject       string `json:"subject"`
	Content       string `json:"content"`
	Options       string `json:"options,omitempty"` // 选择题有值
	Difficulty    int 	 `json:"difficulty"`
	Author        string `json:"author"`
}

func getQuestionsByTypeID(db *gorm.DB, questionType string, ids []int) ([]QuestionSummary, error) {
	var questions []QuestionSummary

	var tableName string
	switch questionType {
	case "Choice":
		tableName = "choice_questions"
	case "Subjective":
		tableName = "subjective_questions"
	default:
		return nil, fmt.Errorf("unsupported question type: %s", questionType)
	}

	// 查询并映射到QuestionSummary
	if err := db.Table(tableName).Select("id, 'Choice' AS question_type, subject, content, options, difficulty, author").Where("id IN (?)", ids).Scan(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}
