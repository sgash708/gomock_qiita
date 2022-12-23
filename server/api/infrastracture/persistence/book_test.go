package persistence_test

import (
	"errors"
	"server/api/domain/model"
	"testing"

	"gorm.io/gorm"
)

func createBook(p *model.Book) (*model.Book, error) {
	pc := d.Create(p)

	return p, pc.Error
}

func findDeletedBook(id int) (*model.Book, error) {
	// DeletedBook
	var p model.Book
	pc := d.Unscoped().First(&p, id)

	return &p, pc.Error
}

func findBook(id int) (*model.Book, error) {
	var p model.Book
	pc := d.First(&p, id)

	err := pc.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &p, err
}

func TestBookFindAllSuccess(t *testing.T) {
	// Delete All
	tearDownDB()

	// Models
	b := model.Book{
		Name: "hoge_name",
		UUID: "test-uuid",
	}

	// Create
	book, err := createBook(&b)
	if err != nil {
		t.Error(err)
	}

	// TestFindAll
	res, err := di.br.FindAll(di.ctx, di.r)
	if err != nil {
		t.Error(err)
	}

	// Check
	if len(res) == 0 {
		t.Error("bookが取得できていません")
	}
	if res[0].Name != book.Name {
		t.Error("book.Nameとmodel.Nameが一致してません")
	}
}

func TestBookFindAllSuccessWithEmptyData(t *testing.T) {
	// Delete All
	tearDownDB()

	// TestFindAll
	book, err := di.br.FindAll(di.ctx, di.r)
	if err != nil {
		t.Error(err)
	}

	// Check
	if len(book) != 0 {
		t.Error("bookの構造体が返されてます")
	}
}

func TestBookFindAllErrorWithEmptyRepo(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Error("nilを引数としてbookを取得できています")
		}
	}()

	// Delete All
	tearDownDB()

	// TestFindAll
	res, err := di.br.FindAll(di.ctx, nil)
	if err == nil {
		t.Errorf("panicにならず取得ができています。\n詳細: %v", err)
	}

	// Check
	if res != nil {
		t.Error("bookの構造体が返されてます")
	}
}

func TestBookFindByUUIDSuccessWithData(t *testing.T) {
	// Delete All
	tearDownDB()

	// Models
	b := model.Book{
		Name: "hoge_name",
		UUID: "test-uuid",
	}

	// Create
	book, err := createBook(&b)
	if err != nil {
		t.Error(err)
	}

	// TestFindByUUID
	res, err := di.br.FindByUUID(di.ctx, di.r, book.UUID)
	if err != nil {
		t.Error(err)
	}

	// Check
	if res.Name != b.Name {
		t.Error("book.Nameとmodel.Nameが一致してません")
	}

	if len(res.CreatedAt.String()) == 0 {
		t.Error("取得したbookに作成日が入っていません")
	}
}

func TestBookFindByUUIDSuccessWithoutdata(t *testing.T) {
	// Delete All
	tearDownDB()

	// TestFindByUUID
	book, err := di.br.FindByUUID(di.ctx, di.r, "test-uuid")
	if err != nil {
		t.Error(err)
	}

	// Check
	if book != nil {
		t.Error("存在しないbookがuuidで取得できています")
	}
}

func TestBookFindByUUIDErrorWithEmptyRepo(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Error("nilを引数としてbookを取得できています")
		}
	}()

	// Delete All
	tearDownDB()

	// TestFindByUUID
	res, err := di.br.FindByUUID(di.ctx, nil, "")
	if err == nil {
		t.Errorf("panicにならず取得ができています。\n詳細: %v", err)
	}

	// Check
	if res != nil {
		t.Error("bookの構造体が返されてます")
	}
}

// func TestBookCreateSuccessWithData(t *testing.T) {
// 	// Delete All
// 	// b.c. gotest is parallel action
// 	tearDownDB()

// 	// Models
// 	p := model.Book{
// 		CompanyID: 1,
// 		Name:      "hoge_name",
// 	}

// 	// TestCreate
// 	book, err := di.pr.Create(di.ctx, di.r, &p)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// Check
// 	if book.CompanyID == 0 {
// 		t.Error("カンパニーIDが生成できていません")
// 	}
// 	if book.Name == "" {
// 		t.Error("Nameが生成できていません")
// 	}
// 	if len(book.CreatedAt.String()) == 0 {
// 		t.Error("bookが生成できていません")
// 	}
// }

// func TestBookCreateErrorWithoutData(t *testing.T) {
// 	// Delete All
// 	// b.c. gotest is parallel action
// 	tearDownDB()

// 	// TestCreate
// 	pc, err := di.pr.Create(di.ctx, di.r, nil)
// 	if err != gorm.ErrInvalidValue {
// 		t.Error(err)
// 	}
// 	if pc != nil {
// 		t.Error("booksが作成できています")
// 	}
// }

// func TestBookDeleteSuccessWithData(t *testing.T) {
// 	// Delete All
// 	// b.c. gotest is parallel action
// 	tearDownDB()

// 	// Models
// 	isDbg := false
// 	l := model.Layout{
// 		Title:        "testtitle",
// 		CustomDesign: "test_custom",
// 	}
// 	ii := model.InputItem{
// 		Type:            "test_type",
// 		Order:           1,
// 		Key:             "name",
// 		MetaData:        "{'placeholder': '田中太郎', 'label': '氏名'}",
// 		ValidationRules: "[{'type': 1}]",
// 	}
// 	iig := model.InputItemGroup{
// 		Type: "nice_type",
// 		InputItems: []*model.InputItem{
// 			&ii,
// 		},
// 	}
// 	m := model.Message{
// 		Type:           "type",
// 		Order:          1,
// 		ImageURL:       "http://localhost:8080/test",
// 		TextHTML:       "<strong>はじめまして</strong>",
// 		RenderType:     "render type",
// 		ConditionRules: "[{ 'key': 2, 'operator': 'hoge', 'value': 'hoge' }]",
// 	}
// 	si := model.ScenarioItem{
// 		ShowSpeed:      1000,
// 		ConditionRules: "[{ 'input_item_id': 2, 'operator': 'hoge', 'value': 'hoge' }]",
// 		InputItemGroups: []*model.InputItemGroup{
// 			&iig,
// 		},
// 		Messages: []*model.Message{
// 			&m,
// 		},
// 	}
// 	pm := model.PipMovie{
// 		Name:     "test_name",
// 		MovieUrl: "http://localhost/test/",
// 	}
// 	imp := model.Impression{}
// 	cv := model.Conversion{}
// 	el := model.EventLog{}
// 	sess := model.Session{
// 		UUID: "test",
// 		Impressions: []*model.Impression{
// 			&imp,
// 		},
// 		EventLogs: []*model.EventLog{
// 			&el,
// 		},
// 		Conversions: []*model.Conversion{
// 			&cv,
// 		},
// 	}
// 	s := model.Scenario{
// 		Name:    "hogemaru_name",
// 		IsDebug: &isDbg,
// 		Layout:  &l,
// 		ScenarioItems: []*model.ScenarioItem{
// 			&si,
// 		},
// 		PipMovies: []*model.PipMovie{
// 			&pm,
// 		},
// 		Sessions: []*model.Session{
// 			&sess,
// 		},
// 	}
// 	p := model.Book{
// 		CompanyID: 1,
// 		Name:      "hoge_name",
// 		Scenarios: []*model.Scenario{
// 			&s,
// 		},
// 	}

// 	// Create
// 	book, err := createBook(&p)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// TestDelete
// 	if err := di.pr.Delete(di.ctx, di.r, book); err != nil {
// 		t.Error(err)
// 	}

// 	// FindByID (Deleted)
// 	res, err := findDeletedBook(book.ID)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// Check
// 	if len(res.DeletedAt.Time.String()) == 0 {
// 		t.Error("削除したbooksに削除日が入っていません")
// 	}

// 	// FindByID
// 	res, err = findBook(book.ID)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	// first level of nesting
// 	unscopeScenario, err := findLayout(book.Scenarios[0].ID)
// 	if err != nil || unscopeScenario != nil {
// 		t.Error("scenariosが削除できていません")
// 	}
// 	// second level of nesting
// 	unscopeLayout, err := findLayout(book.Scenarios[0].Layout.ID)
// 	if err != nil || unscopeLayout != nil {
// 		t.Error("layoutsが削除できていません")
// 	}
// 	// third level of nesting
// 	unscopeInputItemGroup, err := findInputItemGroup(book.Scenarios[0].ScenarioItems[0].InputItemGroups[0].ID)
// 	if err != nil || unscopeInputItemGroup != nil {
// 		t.Error("input_item_groupsが削除できていません")
// 	}
// 	// fourth level of nesting
// 	unscopeInputItem, err := findInputItem(book.Scenarios[0].ScenarioItems[0].InputItemGroups[0].InputItems[0].ID)
// 	if err != nil || unscopeInputItem != nil {
// 		t.Error("input_itemsが削除できていません")
// 	}

// 	// Check
// 	if res != nil {
// 		t.Error("削除したbooksがid指定して取得できています")
// 	}
// }

// func TestBookDeleteErrorWithoutData(t *testing.T) {
// 	// Delete All
// 	// b.c. gotest is parallel action
// 	tearDownDB()

// 	// TestDelete
// 	err := di.pr.Delete(di.ctx, di.r, nil)
// 	if err != gorm.ErrInvalidValue {
// 		t.Error(err)
// 	}
// }

// func TestBookUpdateSuccessWithData(t *testing.T) {
// 	// Delete All
// 	// b.c. gotest is parallel action
// 	tearDownDB()

// 	// Models
// 	p := model.Book{
// 		CompanyID: 1,
// 		Name:      "hoge_name",
// 	}

// 	// Create
// 	book, err := createBook(&p)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// Update Request Data
// 	reqBook := model.Book{
// 		ID:   book.ID,
// 		Name: "hoge_hoge_name",
// 	}

// 	// TestUpdate
// 	res, err := di.pr.Update(di.ctx, di.r, &reqBook)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// Check
// 	if res.Name == book.Name {
// 		t.Error("bookのNameが更新されていません")
// 	}
// }

// func TestBookUpdateErrorWithoutData(t *testing.T) {
// 	defer func() {
// 		err := recover()
// 		if err == nil {
// 			t.Error("nilを引数としてbookを更新できています")
// 		}
// 	}()

// 	// Delete All
// 	// b.c. gotest is parallel action
// 	tearDownDB()

// 	// TestUpdate
// 	res, err := di.pr.Update(di.ctx, di.r, nil)
// 	if err != nil {
// 		t.Errorf("panicにならず更新がされエラーが発生しています\n詳細: %v", err)
// 	}

// 	// Check
// 	if res != nil {
// 		t.Error("bookの構造体が返されています")
// 	}
// }

// func TestBookUpdateErrorWithNotExistID(t *testing.T) {
// 	// Delete All
// 	// b.c. gotest is parallel action
// 	tearDownDB()

// 	// Models
// 	p := model.Book{
// 		CompanyID: 1,
// 		Name:      "hoge_name",
// 	}

// 	// Create
// 	_, err := createBook(&p)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// Update Request Data
// 	reqBook := model.Book{
// 		ID:   1100,
// 		Name: "hoge_hoge_name",
// 	}

// 	// TestUpdate
// 	_, err = di.pr.Update(di.ctx, di.r, &reqBook)
// 	if err != nil && err != gorm.ErrInvalidValue {
// 		t.Errorf("存在しない値で更新した際に、不正な値以外のエラーが発生しています\n詳細: %v", err)
// 	}
// }
