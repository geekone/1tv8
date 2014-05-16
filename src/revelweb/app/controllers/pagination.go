package controllers

import(
	"strconv"
)

const (
	ROWS_PER_PAGE = 2		//每页几条记录
	PAGES_PER_VIEW = 10		//最多显示几个页码
)

type Pagination struct{
	page 		int 		//当前页码
	rowCount	int			//记录总数
	url			string		//页码连接
	pageCount	int			//总页数 
}


type PageNum struct {
	Num int
	Url string
	IsCurrent bool
}


// page 当前页, rowCount 记录总数  url 连接,返回 pagination对象
func NewPagination(page,rowCount int,url string) *Pagination{
	pageCount := rowCount / ROWS_PER_PAGE		//先得总页数 总记录除每页几条记录数

	//如果总页数与每页记录数总和还小于总记录数,总页数加1
	if pageCount * ROWS_PER_PAGE < rowCount {
		pageCount += 1
	}

	return &Pagination{
		page: page,
		rowCount:rowCount,
		url:url,
		pageCount:pageCount,
	}
}	


func (p *Pagination) Pages() []PageNum {
	var result []PageNum

	if p.pageCount == 1 {
		return result
	}

	// 计算起始位置
	begin := p.page - PAGES_PER_VIEW/2
	if begin < 0 {
		begin = 0
	}

	to := begin + PAGES_PER_VIEW
	if to > p.pageCount {
		to = p.pageCount
	}

	// 确定页码数量
	if to-begin < PAGES_PER_VIEW {
		begin = to - PAGES_PER_VIEW
	}
	if begin < 0 {
		begin = 0
	}

	pageNum := 0
	for ; begin < to; begin++ {
		pageNum = begin + 1
		result = append(result, PageNum{pageNum, p.url + strconv.Itoa(pageNum), pageNum == p.page})
	}

	return result
}
