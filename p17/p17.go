package main

import "fmt"

type Tree struct {
	root *Branch // корень дерева
}

type Branch struct {
	val    int     // значение
	left   *Branch // левая ветка
	right  *Branch // правая ветка
	parent *Branch // родитель ветки
}

// Несбалансированое дерево поиска

func main() {
	tree := NewTree(25)
	tree.Insert(26)
	tree.Insert(24)
	tree.Delete(25)
	tree.Delete(24)
	fmt.Printf("%v", *tree.root) // 26 nil nil
}

func NewTree(startVal int) *Tree { // создаем дерево с корнем, родитель у которого nil
	return &Tree{
		root: &Branch{
			val:    startVal,
			parent: nil,
		},
	}
}

func (t *Tree) Insert(val int) { // вставка в дерево через вставку в корень
	t.root.insert(val)
}

func (b *Branch) insert(val int) {
	if val > b.val { // если значение больше, чем то что в ветке - смотрим направо
		if b.right == nil { // если справа пусто, мы это право создадим, родителем будет ветка, которая вызвана
			b.right = &Branch{
				val:    val,
				parent: b,
			}
			return
		}
		b.right.insert(val) // если справа есть путь дальше, идем направо рекурсивно
		return
	} else if val == b.val { // если значение такое же - выходим
		return
	} else { // если значение меньше, чем то что в ветке - смотрим налево
		if b.left == nil { // если слева пусто, мы это лево создадим, родителем будет ветка, которая вызвана
			b.left = &Branch{
				val:    val,
				parent: b,
			}
			return
		}
		b.left.insert(val) // если соева есть путь дальше, идем налево рекурсивно
	}
}

func (t *Tree) Find(val int) bool { // поиск (возвращает результат, существует ли значение)
	return t.root.Find(val)
}

func (b *Branch) Find(val int) bool {
	if val > b.val { // Если больше - идем направо
		if b.right == nil { // если справа ничего нет - значит, элемента нет в дереве
			return false
		}
		return b.right.Find(val) // если право есть - ищем справа рекурсивно
	} else if val == b.val { // нашли значение, результат true
		return true
	} else { // если меньше - идем налево
		if b.left == nil { // Если слева ничего нет - значит, элемента нет в дереве
			return false
		}
		return b.left.Find(val) // если лево есть - ищем слева рекурсивно
	}
}

func (t *Tree) Delete(val int) bool { // удаляем из дерева с корня
	return t.root.Delete(nil, val)
}

func (b *Branch) Delete(root *Branch, val int) bool {
	if val > b.val { // если значение больше корня - ищем справа
		if b.right == nil { // право не существует, результат удаления false
			return false
		}
		return b.right.Delete(b, val) // ищем справа рекурсивно
	} else if val == b.val { // если результат тот, что мы нашли, то...
		if b.right == nil && b.left == nil { // если детей нет, то...
			if root == nil { // если нет родителя, то ветку можно сделать пустой, это корень, нашли то что нужно
				b = &Branch{}
				return true
			}
			if root.right == b { // если правая ветка родителя это b - чистим её
				root.right = nil
			} else { // если левая ветка родителя это b - чистим её
				root.left = nil
			}
			b.val = 0
		} else if b.right == nil || b.left == nil { // если нет хотя бы 1 ребенка, то
			var notEmpty *Branch // ищем ребенка, который есть
			if b.right != nil {
				notEmpty = b.right
			} else {
				notEmpty = b.left
			}
			if root == nil { // если корня у этого нет, значит это корень и нужно поменять ему значение
				*b = *notEmpty // меняем значение буквально по ссылке
				b.parent = nil // родителя ставим nil
				return true    // возвращаем, что всё нашлось и удалено
			}
			if root.right == b { // если эта ветка - это правая часть родителя, ставим ребенка ребенком родителя справа
				root.right = notEmpty
				notEmpty.parent = root
			} else { // если эта ветка - это левая часть родителя, ставим ребенка ребенком родителя слева
				root.left = notEmpty
				notEmpty.parent = root
			}
		} else {
			successor := b.right.Next()             // ищем следующее значение (минимально большое после текущего)
			b.val = successor.val                   // ставим текущим значением следующее
			if successor.parent.left == successor { // если следующее значение находится слева от родителя
				successor.parent.left = successor.right // делаем левое значение родителя правым значением следующего
				if successor.right != nil {             // если правое значение следующего не пустое
					successor.right.parent = successor.parent // ставим родителем правого значения родителя следующего
				}
			} else { // если следующее значение находится слева от родителя
				successor.parent.right = successor.right // то правое значение родителя теперь правое значение следующего
				if successor.right != nil {
					successor.right.parent = successor.parent // ставим родителем правого значения родителя следующего
				}
			}
		}
		return true // мы нашли и все сделали - значение удалено
	} else {
		if b.left == nil { // если слева пусто, то удалять нечего, false
			return false
		}
		return b.left.Delete(b, val) // если слева есть путь - идем по нему рекурсивно
	}
}

func (b *Branch) Next() *Branch { // ищем следующее минимальнейшее значение от указаной ветки
	if b.left != nil { // если ветка слева есть
		return b.left.Next() // идем по ней
	}
	return b // если ветки слева нет, возвращаем текущую как самую минимальную, потому что справа только больше
}
