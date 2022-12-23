package persistence_test

import (
	"errors"
	"server/api/domain/model"
	"testing"

	"gorm.io/gorm"
)

func createBook(b *model.Book) (*model.Book, error) {
	book := d.Create(b)

	return b, book.Error
}

func findDeletedBook(id int) (*model.Book, error) {
	var b model.Book
	book := d.Unscoped().First(&b, id)

	return &b, book.Error
}

func findBook(id int) (*model.Book, error) {
	var b model.Book

	err := d.First(&b, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &b, err
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

func TestBookCreateSuccessWithData(t *testing.T) {
	// Delete All
	tearDownDB()

	// Models
	b := model.Book{
		UUID: "test-uuid",
		Name: "hoge_name",
	}

	// TestCreate
	book, err := di.br.Create(di.ctx, di.r, &b)
	if err != nil {
		t.Error(err)
	}

	// Check
	if book.Name == "" {
		t.Error("Nameが生成できていません")
	}
	if len(book.CreatedAt.String()) == 0 {
		t.Error("bookが生成できていません")
	}
}

func TestBookCreateErrorWithoutData(t *testing.T) {
	// Delete All
	tearDownDB()

	// TestCreate
	b, err := di.br.Create(di.ctx, di.r, nil)
	if err != gorm.ErrInvalidValue {
		t.Error(err)
	}
	if b != nil {
		t.Error("bookが作成できています")
	}
}

func TestBookUpdateSuccessWithData(t *testing.T) {
	// Delete All
	tearDownDB()

	// Models
	b := model.Book{
		UUID: "test-uuid",
		Name: "hoge_name",
	}

	// Create
	book, err := createBook(&b)
	if err != nil {
		t.Error(err)
	}

	// Update Request Data
	reqBook := model.Book{
		ID:   book.ID,
		UUID: book.UUID,
		Name: "hoge_hoge_name",
	}

	// TestUpdate
	res, err := di.br.Update(di.ctx, di.r, &reqBook)
	if err != nil {
		t.Error(err)
	}

	// Check
	if res.Name == book.Name {
		t.Error("bookのNameが更新されていません")
	}
}

func TestBookUpdateErrorWithoutData(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Error("nilを引数としてbookを更新できています")
		}
	}()

	// Delete All
	tearDownDB()

	// TestUpdate
	res, err := di.br.Update(di.ctx, di.r, nil)
	if err != nil {
		t.Errorf("panicにならず更新がされエラーが発生しています\n詳細: %v", err)
	}

	// Check
	if res != nil {
		t.Error("bookの構造体が返されています")
	}
}

func TestBookDeleteSuccessWithData(t *testing.T) {
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

	// TestDelete
	if err := di.br.Delete(di.ctx, di.r, book); err != nil {
		t.Error(err)
	}

	// FindByID (Deleted)
	res, err := findDeletedBook(book.ID)
	if err != nil {
		t.Error(err)
	}

	// Check
	if len(res.DeletedAt.Time.String()) == 0 {
		t.Error("削除したbooksに削除日が入っていません")
	}

	// FindByID
	res, err = findBook(book.ID)
	if err != nil {
		t.Error(err)
	}

	// Check
	if res != nil {
		t.Error("削除したbooksがid指定して取得できています")
	}
}

func TestBookDeleteErrorWithoutData(t *testing.T) {
	// Delete All
	tearDownDB()

	// TestDelete
	if err := di.br.Delete(di.ctx, di.r, nil); err != gorm.ErrInvalidValue {
		t.Error(err)
	}
}
