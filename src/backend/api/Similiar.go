package api

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"log"
)
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
func Findsimiliar(c *gin.Context) {
	var inputIDs []int
	if err := c.ShouldBindJSON(&inputIDs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "reason": "Invalid input format"})
		return
	}

	db, err := getDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Failed to connect to database"})
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

// addEdge 添加有向边u->v，容量cap，费用cost
func (g *Graph) addEdge(u, v, cap, cost int) {
	g.List[u] = append(g.List[u], Edge{V: v, Cap: cap, Cost: cost, Rev: len(g.List[v])})
	g.List[v] = append(g.List[v], Edge{V: u, Cap: 0, Cost: -cost, Rev: len(g.List[u])-1}) // 反向边
}


func MaxFlow(db *gorm.DB, choiceIDs []int, distinctChoiceIDs []int) {
	var num int
	for _, fromID := range choiceIDs {
		num=max(num,fromID)
	}
	for _, toID := range distinctChoiceIDs {
		num=max(num,toID)
	}

	g := Graph{N: num+2} 
	s := num+1 // 源点
	t := num+2 // 汇点

	for _, fromID := range choiceIDs {
		g.addEdge(s, fromID, 1, 0)
	}
	for _, toID := range distinctChoiceIDs {
		g.addEdge(toID, t, 1, 0)
	}

	questionKeywordsMap := GetAllKeywords(db, append(choiceIDs, distinctChoiceIDs...))

	for _, fromID := range choiceIDs {
		for _, toID := range distinctChoiceIDs {
			go func(fid, tid int) {
				defer wg.Done()
				similarity := calculateSharedKeywordCount(questionKeywordsMap[fid], questionKeywordsMap[tid])
				if similarity > 0 {
					g.addEdge(fid,tid,1,similarity+5)
				}
				else{
					g.addEdge(fid,tid,1,1)
				}
			}(fromID, toID)
		}
	}

	flow, cost := g.dinic(s, t)
	fmt.Printf("最大流: %d, 最小费用: %d\n", flow, cost)

}

func (g *Graph) dinicDfs(v, t, f int, level []int, iter []int) int {
	if v == t {
		return f
	}
	for i := iter[v]; i < len(g.List[v]); i++ {
		e := &g.List[v][i]
		if e.Cap > 0 && level[v] < level[e.V] {
			d := g.dinicDfs(e.V, t, min(f, e.Cap), level, iter)
			if d > 0 {
				e.Cap -= d
				g.List[e.V][e.Rev].Cap += d
				return d
			}
		}
	}
	return 0
}

func (g *Graph) dinic(s, t int) (int, int) {
	flow, cost := 0, 0
	for {
		level := make([]int, g.N)
		iter := make([]int, g.N)
		maxLevel := g.bfs(s, t, level)
		if maxLevel == -1 {
			break 
		}
		for {
			f := g.dinicDfs(s, t, 1<<60, level, iter)
			if f == 0 {
				break
			}
			flow += f
			// 累加沿路径的费用
		}
	}
	return flow, cost
}

// bfs 求解层次图
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
