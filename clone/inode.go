package main

// 这个interface设计得可以用在文件和文件夹
// 所有的原型类都实现了接口的clone方法
// 但如果file是read，folder是open呢
type inode interface {
	print(string)
	clone() inode
}
