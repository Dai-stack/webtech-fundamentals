package main

import (
	"html/template"
	"log"
	"net/http"
)

var todoList []string // <1>

func handleTodo(w http.ResponseWriter, r *http.Request) { // <5>
	t, _ := template.ParseFiles("templates/todo.html") // <6> テンプレートを解析し、変数にいれる
	// テンプレートと変数toDoListを組み合わせ動的にHTMLを生成し、レスポンスとして返している。
	t.Execute(w, todoList)
}

func main() {
	todoList = append(todoList, "顔を洗う", "朝食を食べる", "歯を磨く") // <2>

	// １つ目のstatic→クライアントがアクセ　するURLパス　ファイルシステム上で静的ファイルを探すための基点
	// 2つ目のstatic→サーバー側のファイルが存在するディレクトリ名
	http.Handle("/static/", //
		http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // <3>

	//　パスtodoでアクセスがきた場合、関数handleTodoを返す
	http.HandleFunc("/todo", handleTodo) // <4>

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start : ", err)
	}
}
