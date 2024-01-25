package mall

import (
	"fmt"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/mall"
	mallRes "main.go/model/mall/response"
	"main.go/model/manage"
	"main.go/utils"
	"strconv"
	"strings"
)

type MallQueryService struct {
}

func (m *MallQueryService) TextToItem(text string) (err error, cartItems []mallRes.RecommendResponse) {
	embedding, err := utils.GetEmbedding(text)
	if err != nil {
		return err, nil
	}
	fmt.Println(len(embedding))
	ids, err := utils.PgvectorGetId(embedding)
	fmt.Println(ids)
	var resItems []manage.MallGoodsInfo
	err = global.GVA_DB.Where("goods_id IN (?)", ids).Find(&resItems).Error
	if err != nil {
		return err, nil
	}
	for _, items := range resItems {
		cartItems = append(cartItems, mallRes.RecommendResponse{
			GoodsId:       items.GoodsId,
			GoodsName:     items.GoodsName,
			GoodsCoverImg: items.GoodsCoverImg,
			SellingPrice:  items.SellingPrice,
		})
	}
	return
}

type ProductRecommend struct {
	Score int
	Items []mallRes.RecommendResponse
}

type ProductScore struct {
	Name  string
	Score int
}

func ParseProductScores(input string) ([]ProductScore, error) {
	var products []ProductScore
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid format: %s", line)
		}
		score, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid score format: %s", parts[1])
		}
		products = append(products, ProductScore{
			Name:  strings.TrimSpace(parts[0]),
			Score: score,
		})
	}
	return products, nil
}

func (m *MallQueryService) FineTuneGetList(text string) (err error, cartItems []mallRes.RecommendResponse) {
	userid := 7
	var products []ProductScore
	for i := 0; i < 3; i++ {
		openAIRes, err := utils.CallOpenAI(text)
		if err == nil {
			//fmt.Println(openAIRes)
			products, err = ParseProductScores(openAIRes)
			if err != nil {
				continue
			}
			break
		}
		global.GVA_LOG.Error("  Attempt "+strconv.FormatInt(int64(i+1), 10)+" times failed: %v\n", zap.Error(err))
	}
	//fmt.Println("ParseProductScores Res:", products)
	var ModelInput []ProductRecommend
	for _, items := range products {
		err, cartItems := m.TextToItem(items.Name)
		if err != nil {
			global.GVA_LOG.Error("获取推荐失败", zap.Error(err))
		}
		ModelInput = append(ModelInput, ProductRecommend{Items: cartItems,
			Score: items.Score,
		})
	}
	fmt.Println("Model final input=", ModelInput)
	var userFinance mall.MallUserFinance
	err = global.GVA_DB.Where("user_id = ? ", userid).First(&userFinance).Error
	if err != nil {
		return err, nil
	}
	for _, items := range ModelInput {
		cartItems = append(cartItems, items.Items[0])
	}
	return
}
