package api

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindSamePost(c *gin.Context) {
	var form struct {
		Username string `form:"username" binding:"required"`
		TestId   int    `form:"testId" binding:"required"`
	}

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
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users not found"})
		return
	}

	// 鉴权，学生无法操作
	if user.Type == STUDENT {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Permission denied"})
		return

	}
	//获取id列表
	inputIDs := queryTestByID(db, form.TestId)

	//分割
	var choiceIDs []int
	var subjectiveIDs []int

	for _, id := range inputIDs {
		questionType, _ := getQuestionTypeByID(db, id)
		switch questionType {
		case "choice":
			choiceIDs = append(choiceIDs, id)
		case "subjective":
			subjectiveIDs = append(subjectiveIDs, id)
		default:
			log.Printf("Unknown question type for ID %d", id)
		}
	}
	log.Println("begin searching")
	// 查询不重复的选择题ID
	var distinctChoiceIDs []int
	if err := db.Table("choicequestions").Select("id").Where("id NOT IN ?", append(choiceIDs, 0)).Scan(&distinctChoiceIDs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Failed to fetch distinct choice question IDs"})
		return
	}

	// 查询不重复的主观题ID
	var distinctSubjectiveIDs []int
	if err := db.Table("subjectivequestions").Select("id").Where("id NOT IN ?", append(subjectiveIDs, 0)).Scan(&distinctSubjectiveIDs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Failed to fetch distinct subjective question IDs"})
		return
	}

	ChoicesNodes := MaxFlow(db, choiceIDs, distinctChoiceIDs, "choice")
	SubjectiveNodes := MaxFlow(db, subjectiveIDs, distinctSubjectiveIDs, "subjective")

	choicesQuestion, _ := getQuestionsByTypeID(db, "Choice", ChoicesNodes)
	subjectiveQuestion, _ := getQuestionsByTypeID(db, "Subjective", SubjectiveNodes)

	var questions []QuestionSummary
	questions = append(questions, choicesQuestion...)
	questions = append(questions, subjectiveQuestion...)

	log.Println("finish searching")
	var retQuestions []QuestionSummary
	for _, ques := range questions {
		var temp QuestionSummary
		temp.ID = ques.ID
		if ques.QuestionType == "choicequestions" {
			temp.QuestionType = "choice_questions"
		} else {
			temp.QuestionType = "subjective_questions"
		}
		temp.Subject = ques.Subject
		temp.Content = ques.Content
		temp.Difficulty = ques.Difficulty
		temp.Author = ques.Author
		retQuestions = append(retQuestions, temp)

	}
	// fmt.Println(len(retQuestions))
	if len(retQuestions) == 0 {
		c.JSON(http.StatusOK, gin.H{"success": false, "reason": "No questions found"})
		return
	}
	mdFile, _ := GenerateMdByQuestions(db, retQuestions)
	c.JSON(http.StatusOK, gin.H{"success": true, "test": mdFile})
}

func GetAllKeywords(db *gorm.DB, questionIDs []int, questionType string) map[int][]string {
	questionKeywordsMap := make(map[int][]string)
	for _, qID := range questionIDs {
		var keywords []string
		if questionType == "choice" {
			db.Table("choicequestions").Joins("JOIN choice_question_keywords ON choicequestions.id = choice_question_keywords.question_id").
				Joins("JOIN keywords ON choice_question_keywords.keyword_id = keywords.id").
				Where("choicequestions.id = ?", qID).
				Pluck("keywords.keyword", &keywords)
		} else {
			db.Table("subjectivequestions").Joins("JOIN subjective_question_keywords ON subjectivequestions.id = subjective_question_keywords.question_id").
				Joins("JOIN keywords ON subjective_question_keywords.keyword_id = keywords.id").
				Where("subjectivequestions.id = ?", qID).
				Pluck("keywords.keyword", &keywords)
		}
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
	V    int // 目标顶点
	Cap  int // 剩余容量
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
	// fmt.Println("edge", u, v, cap, cost)
	g.List[u] = append(g.List[u], Edge{V: v, Cap: cap, Cost: cost, Rev: len(g.List[v])})
	g.List[v] = append(g.List[v], Edge{V: u, Cap: 0, Cost: -cost, Rev: len(g.List[u]) - 1}) // 反向边
}

func MaxFlow(db *gorm.DB, IDs []int, distinctIDs []int, questionType string) []int {
	var num int
	for _, fromID := range IDs {
		num = mmax(num, fromID)
	}
	for _, toID := range distinctIDs {
		num = mmax(num, toID)
	}

	g := NewGraph(num + 2)
	s := num + 1 // 源点
	t := num + 2 // 汇点

	for _, fromID := range IDs {
		g.addEdge(s, fromID, 1, 0)
	}
	// fmt.Println(IDs)
	for _, toID := range distinctIDs {
		g.addEdge(toID, t, 1, 0)
	}
	// fmt.Println(distinctIDs)
	questionKeywordsMap := GetAllKeywords(db, append(IDs, distinctIDs...), questionType)

	var wg sync.WaitGroup

	for _, fromID := range IDs {
		for _, toID := range distinctIDs {
			wg.Add(1)
			go func(fid, tid int) {
				defer wg.Done()
				similarity := calculateSharedKeywordCount(questionKeywordsMap[fid], questionKeywordsMap[tid])

				if similarity > 0 {
					g.addEdge(fid, tid, 1, similarity+5)
				} else {
					g.addEdge(fid, tid, 1, 1)
				}
			}(fromID, toID)
		}
	}

	// 等待所有加边goroutine完成
	wg.Wait()

	// fmt.Println(s, t)
	_, _ = g.dinic(s, t)
	// fmt.Printf("最大流: %d, 最小费用: %d\n", flow, cost)
	rightnodes := g.RightNodes(t)
	return rightnodes
}

func (g *Graph) dinicDfs(v, t, f int, level []int, iter []int, currentCost int) (int, int) {
	if v == t {
		return f, currentCost * f // 累加当前路径的费用
	}
	for i := iter[v]; i < len(g.List[v]); i++ {
		e := &g.List[v][i]
		if e.Cap > 0 && level[v] < level[e.V] {
			d, pathCost := g.dinicDfs(e.V, t, mmin(f, e.Cap), level, iter, currentCost+e.Cost)
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
func (g *Graph) RightNodes(t int) []int {
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

func mmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func mmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
