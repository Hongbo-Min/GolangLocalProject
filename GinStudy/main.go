package main

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/gorilla/sessions"
)

// type Login struct {
// 	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
// 	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
// }

var store = sessions.NewCookieStore([]byte("min-hong-bo"))

// func main() {
// 	// 1. 创建路由
// 	// 2. 默认使用了2个中间件Logger(), Recovery()
// 	r := gin.Default()
// 	// r.Use(MiddleWare())
// 	r.Use(myTime)
// 	// 限制上传文件大小，默认32MB ，当前为8MB
// 	r.MaxMultipartMemory = 8 << 20

// 	v1 := r.Group("/v1")

// 	{
// 		// 路由组1，处理Get请求
// 		v1.GET("/login", login)
// 		v1.GET("/submit", submit)
// 	}

// 	v2 := r.Group("/v2")
// 	{
// 		// 路由组2，处理Post请求
// 		v2.POST("/login", login)
// 		v2.POST("/submit", submit)
// 	}

// 	r.GET("/user", func(c *gin.Context) {
// 		name := c.DefaultQuery("name", "MinHongbo")
// 		c.String(http.StatusOK, "hello "+name)
// 	})
// 	r.POST("/form", func(c *gin.Context) {
// 		types := c.DefaultPostForm("type", "post")
// 		username := c.PostForm("username")
// 		password := c.PostForm("userpassword")
// 		c.String(http.StatusOK, fmt.Sprintf("username:%s, password:%s, type:%s", username, password, types))
// 	})
// 	r.POST("/upload", func(c *gin.Context) {
// 		form, err := c.MultipartForm()
// 		if err != nil {
// 			c.String(http.StatusInternalServerError, fmt.Sprintf("get err %s", err.Error()))
// 		}
// 		files := form.File["files"]
// 		for _, file := range files {
// 			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
// 				c.String(http.StatusInternalServerError, fmt.Sprintf("upload err %s", err.Error()))
// 				return
// 			}
// 		}
// 		c.String(http.StatusOK, fmt.Sprintf("upload ok %d files", len(files)))
// 	})
// 	r.POST("/loginJSON", func(c *gin.Context) {
// 		var json Login
// 		if err := c.ShouldBindJSON(&json); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if json.User != "root" || json.Password != "admin" {
// 			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"status": "200"})
// 	})
// 	r.POST("/loginForm", func(c *gin.Context) {
// 		var form Login
// 		if err := c.Bind(&form); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		if form.User != "root" || form.Password != "admin" {
// 			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"status": "200"})
// 	})
// 	r.GET("/:user/:password", func(c *gin.Context) {
// 		var login Login
// 		if err := c.ShouldBindUri(&login); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		if login.User != "root" || login.Password != "admin" {
// 			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"status": "200"})
// 	})
// 	// json响应
// 	r.GET("/someJSON", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
// 	})
// 	// 结构体响应
// 	r.GET("/someStruct", func(c *gin.Context) {
// 		var msg struct {
// 			Name    string
// 			Message string
// 			Number  int
// 		}
// 		msg.Name = "root"
// 		msg.Message = "message"
// 		msg.Number = 123
// 		c.JSON(200, msg)
// 	})
// 	// XML响应
// 	r.GET("/someXML", func(c *gin.Context) {
// 		c.XML(200, gin.H{"message": "abc"})
// 	})
// 	// YAML响应
// 	r.GET("/someYAML", func(c *gin.Context) {
// 		c.YAML(200, gin.H{"name": "zhangsan"})
// 	})
// 	r.GET("/someProtoBuf", func(c *gin.Context) {
// 		reps := []int64{int64(1), int64(2)}
// 		label := "label"

// 		data := &protoexample.Test{
// 			Label: &label,
// 			Reps:  reps,
// 		}

// 		c.ProtoBuf(200, data)
// 	})

// 	/* r.LoadHTMLGlob("tem/*")
// 	r.GET("/index", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
// 	})
// 	*/
// 	r.GET("/index", func(c *gin.Context) {
// 		c.Redirect(http.StatusMovedPermanently, "http://www.51mh.com")
// 	})
// 	// gin框架实现404页面
// 	r.NoRoute(func(c *gin.Context) {
// 		c.String(http.StatusNotFound, "404 not found 22222")
// 	})

// 	// 异步
// 	r.GET("/long_async", func(c *gin.Context) {
// 		copyContext := c.Copy()

// 		go func() {
// 			time.Sleep(3 * time.Second)
// 			log.Println("异步执行: " + copyContext.Request.URL.Path)
// 		}()
// 	})
// 	// 同步
// 	r.GET("/long_sync", func(c *gin.Context) {
// 		time.Sleep(3 * time.Second)
// 		log.Println("同步执行: " + c.Request.URL.Path)
// 	})

// 	/* r.GET("/ce", func(c *gin.Context) {
// 		req, _ := c.Get("request")
// 		fmt.Println("request: ", req)
// 		c.JSON(200, gin.H{"request": req})
// 	}) */
// 	r.GET("/ce", MiddleWare(), func(c *gin.Context) {
// 		req, _ := c.Get("request")
// 		fmt.Println("request: ", req)
// 		c.JSON(200, gin.H{"request": req})
// 	})
// 	shoppingGroup := r.Group("/shopping")
// 	{
// 		shoppingGroup.GET("/index", shopIndexHandler)
// 		shoppingGroup.GET("/home", shopHomeHandler)
// 	}

// 	r.GET("/cookie", func(c *gin.Context) {
// 		cookie, err := c.Cookie("key_cookie")
// 		if err != nil {
// 			cookie = "NotSet"
// 			c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
// 		}
// 		fmt.Printf("cookie的值是: %s\n", cookie)
// 	})

// 	r.GET("/login", func(c *gin.Context) {
// 		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
// 		c.String(200, "Login success!")
// 	})

// 	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
// 		c.JSON(200, gin.H{"data": "home"})
// 	})

// 	r.Run(":8000")
// }

// func main() {
// 	http.HandleFunc("/save", SaveSession)
// 	http.HandleFunc("/get", GetSession)

// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		fmt.Println("HTTP server failed, err: ", err)
// 		return
// 	}
// }

type Person struct {
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

// func main() {
// 	r := gin.Default()
// 	r.GET("/5lmh", func(c *gin.Context) {
// 		var person Person
// 		if err := c.ShouldBind(&person); err != nil {
// 			c.String(500, fmt.Sprint(err))
// 			return
// 		}
// 		c.String(200, fmt.Sprintf("%#v", person))
// 	})
// 	r.Run()
// }

type Login struct {
	User     string `uri:"user" validate:"checkName"`
	Password string `uri:"password"`
}

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" timme_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	//field.Interface().(time.Time)获取参数值并且转换为时间格式
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Unix() > date.Unix() {
			return false
		}
	}
	return true
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func checkName(fl validator.FieldLevel) bool {
	return fl.Field().String() == "root"
}

func main() {
	r := gin.Default()
	validate := validator.New()
	r.GET("/:user/:password", func(c *gin.Context) {
		var login Login
		err := validate.RegisterValidation("checkName", checkName)
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = validate.Struct(login)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err)
			}
			return
		}
		fmt.Println("success")
	})

	r.GET("/5lmh", getBookable)
	r.Run()
}

func SaveSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "min")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	session.Save(r, w)
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "min")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	foo := session.Values["foo"]
	fmt.Println("foo: ", foo)
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		c.Set("request", "中间件")
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time: ", t2)
	}
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		c.Abort()
		return
	}
}

func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	since := time.Since(start)
	fmt.Println("程序用时: ", since)
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}
