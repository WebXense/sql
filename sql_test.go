package sql_test

import (
	"os"

	"github.com/WebXense/sql"
	"github.com/WebXense/sql/conn"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

var _ = Describe("Test 'sql'", func() {

	var db *gorm.DB
	var testDBPath = "test.db"

	BeforeEach(func() {
		var err error
		db, err = conn.SQLite(testDBPath, true)
		Expect(err).To(BeNil())

		db.AutoMigrate(&User{})

		users := []User{
			{Name: "John", Age: 18},
			{Name: "Jane", Age: 20},
			{Name: "Jack", Age: 22},
		}
		err = db.Create(&users).Error
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		os.RemoveAll(testDBPath)
	})

	It("should be able to create a record", func() {
		user := &User{Name: "Peter", Age: 25}
		var err error
		user, err = sql.Create(db, user)
		Expect(err).To(BeNil())
		Expect(user.ID).NotTo(BeZero())
	})

	It("should be able to update a record", func() {
		user, err := sql.FindOne[User](db, sql.Eq("name", "John"))
		Expect(err).To(BeNil())
		user.Age = 19
		user, err = sql.Update(db, user)
		Expect(err).To(BeNil())
		Expect(user.Age).To(Equal(19))

		user, err = sql.FindOne[User](db, sql.Eq("name", "John"))
		Expect(err).To(BeNil())
		Expect(user.Age).To(Equal(19))
	})

	It("should be able to delete a record", func() {
		user, err := sql.FindOne[User](db, sql.Eq("name", "John"))
		Expect(err).To(BeNil())

		err = sql.Delete(db, user)
		Expect(err).To(BeNil())

		user, _ = sql.FindOne[User](db, sql.Eq("name", "John"))
		Expect(user).To(BeNil())
	})

	It("should be able to delete multiple records by condition", func() {
		err := sql.DeleteBy[User](db, sql.Lte("age", 20))
		Expect(err).To(BeNil())

		users, err := sql.FindAll[User](db, nil, nil, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(1))
	})

	It("should be able to find record using condition 'eq'", func() {
		user, err := sql.FindOne[User](db, sql.Eq("name", "John"))
		Expect(err).To(BeNil())
		Expect(user.Name).To(Equal("John"))
	})

	It("should be able to find record using condition 'neq'", func() {
		users, err := sql.FindAll[User](db, sql.Neq("name", "John"), nil, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(2))
	})

	It("should be able to find record using condition 'gt'", func() {
		users, err := sql.FindAll[User](db, sql.Gt("age", 20), nil, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(1))
	})

	It("should be able to find record using condition 'gte'", func() {
		users, err := sql.FindAll[User](db, sql.Gte("age", 20), nil, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(2))
	})

	It("should be able to find record using condition 'lt'", func() {
		users, err := sql.FindAll[User](db, sql.Lt("age", 20), nil, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(1))
	})

	It("should be able to find record using condition 'lte'", func() {
		users, err := sql.FindAll[User](db, sql.Lte("age", 20), nil, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(2))
	})

	It("should be able to find record using condition 'in'", func() {
		users, err := sql.FindAll[User](db, sql.In("name", []string{"John", "Jane"}), nil, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(2))
	})

	It("should be able to find record using condition 'nin'", func() {
		users, err := sql.FindAll[User](db, sql.Nin("name", []string{"John", "Jane"}), nil, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(1))
	})

	It("should be able to find record using condition 'lk'", func() {
		users, err := sql.FindAll[User](db, sql.Lk("name", "Ja%"), nil, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(2))
	})

	It("should be able to find record using condition 'nlk'", func() {
		users, err := sql.FindAll[User](db, sql.Nlk("name", "Ja%"), nil, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(1))
	})

	It("should be able to find record using condition 'null'", func() {
		users, err := sql.FindAll[User](db, sql.Null("name"), nil, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(0))
	})

	It("should be able to find record using condition 'not null'", func() {
		users, err := sql.FindAll[User](db, sql.NotNull("name"), nil, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(3))
	})

	It("should be able to find record using condition 'and'", func() {
		users, err := sql.FindAll[User](db, sql.And(sql.In("name", []string{"John", "Jane"}), sql.Eq("age", 20)), nil, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(1))
	})

	It("should be able to find record using condition 'or'", func() {
		users, err := sql.FindAll[User](db, sql.Or(sql.In("name", []string{"John", "Jane"}), sql.Eq("age", 22)), nil, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(3))
	})

	It("should be able to count record using condition", func() {
		count, err := sql.Count[User](db, sql.Gte("age", 20))
		Expect(err).To(BeNil())
		Expect(count).To(Equal(int64(2)))
	})

	It("should be able to find records using pagination", func() {
		users, err := sql.FindAll[User](db, nil, &sql.Pagination{Page: 1, Size: 2}, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(2))

		users, err = sql.FindAll[User](db, nil, &sql.Pagination{Page: 2, Size: 2}, nil)
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(1))
	})

	It("should be able to find records using order", func() {
		users, err := sql.FindAll[User](db, nil, nil, &sql.Sort{SortBy: "age", Asc: true})
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(3))
		Expect(users[0].Age).To(Equal(18))
		Expect(users[1].Age).To(Equal(20))
		Expect(users[2].Age).To(Equal(22))

		users, err = sql.FindAll[User](db, nil, nil, &sql.Sort{SortBy: "age", Asc: false})
		Expect(err).To(BeNil())
		Expect(users).To(HaveLen(3))
		Expect(users[0].Age).To(Equal(22))
		Expect(users[1].Age).To(Equal(20))
		Expect(users[2].Age).To(Equal(18))
	})
})
