package douban

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//var data = []byte(`{"rating":{"max":10,"numRaters":7,"average":"0.0","min":0},"subtitle":"车辆装置篇","author":["[日]五十川芳仁"],"pubdate":"2016-9-1","tags":[{"count":5,"name":"乐高","title":"乐高"},{"count":5,"name":"Lego","title":"Lego"},{"count":3,"name":"工具书","title":"工具书"},{"count":2,"name":"五十川芳仁","title":"五十川芳仁"},{"count":1,"name":"日本","title":"日本"},{"count":1,"name":"基础格栅","title":"基础格栅"},{"count":1,"name":"动力组","title":"动力组"}],"origin_title":"","image":"https://img9.doubanio.com\/view\/subject\/m\/public\/s29086928.jpg","binding":"","translator":["韦皓文"],"catalog":"前言\n第一部分  车辆\n用一个马达驱动车轮\n四轮驱动\n用两个马达驱动车轮\n万向轮\n用伺服马达控制转向\n差速器\n履带车\n小车带动的旋转\n小车带动的移动\n有悬挂系统的小车\n一辆小车与五个不同的底座\n往复运动车\n酷炫小车\n第二部分  没有轮胎的移动\n双足移动\n四足移动\n六足移动\n奇妙的步行者\n尺蠖运动\n振动前进\n其它移动方式\n第三部分  特殊机构\n间歇运动\n转速平稳变化\n切换旋转方向\n用开关完成运动转换\n变速箱\n用旋转方向完成运动转换\n零件表","pages":"318","images":{"small":"https://img9.doubanio.com\/view\/subject\/s\/public\/s29086928.jpg","large":"https://img9.doubanio.com\/view\/subject\/l\/public\/s29086928.jpg","medium":"https://img9.doubanio.com\/view\/subject\/m\/public\/s29086928.jpg"},"alt":"https:\/\/book.douban.com\/subject\/26889294\/","id":"26889294","publisher":"人民邮电出版社","isbn10":"7115431353","isbn13":"9787115431356","title":"乐高动力组创意搭建指南","url":"https:\/\/api.douban.com\/v2\/book\/26889294","alt_title":"","author_intro":"五十川芳仁(Yoshihito Isoqawa)，是一位有46年搭建经验的知名乐高玩家。除了主持五十川乐高工作室之外，他还定期在学校和科学博物馆举办乐高研讨会和讲座，并为展览会和大型活动搭建乐高模型。他同时也是《乐高机器人EV3创意搭建指南》一书的作者。","summary":"本书为五十川芳仁编著的《乐高动力组创意搭建指南(车辆装置篇爱上乐高)》，与《乐高动力组创意搭建指南：机械机构篇》以及《乐高机器人EV3创意搭建指南》为同一系列的书，作者均为国际乐高圈影响力颇高的乐高大师五十川芳仁。本书列出每一个模型的搭建零件，但没有给出搭建步骤，而是给了大量多角度的高清照片，每组照片展示了一个机械原理或一项搭建效果，读者可以观察这些从不同角度拍摄的照片，这种表现方式就像是在解谜，并可尝试重新搭建模型。","series":{"id":"34056","title":"爱上乐高"},"price":"89.00元"}`)

func TestNew(t *testing.T) {
	barcode := "9787115431356"
	actual := New(barcode)
	expected := DoubanBook{Title: "乐高动力组创意搭建指南", Isbn13: barcode}
	assert.Equal(t, expected.Title, actual.Title)
	assert.Equal(t, expected.Isbn13, actual.Isbn13)
}
