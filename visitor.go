package main
 
import "fmt"
 
type visitor interface {
    visitLiquor(*liquor)
    visitTobacco(*tobacco)
    visitNecessity(*necessity)
}
 
type item interface {
    getType() string
    accept(visitor)
}
 
// Liquor inits
type liquor struct {
    price float32
}
 
func (l *liquor) accept(v visitor) {
    v.visitLiquor(l)
}
 
// tobacco inits
type tobacco struct {
    price float32
}
 
func (t *tobacco) accept(v visitor) {
    v.visitTobacco(t)
}
 
// necessity inits
type necessity struct {
    price float32
}
 
func (n *necessity) accept(v visitor) {
    v.visitNecessity(n)
}
 
// concrete type for tax calc behavior
type taxCalculator struct {
    value float32
}
 
func (tx *taxCalculator) visitLiquor(l *liquor) {
    fmt.Printf("Price of liquor item is = $%v \n", l.price)
}
 
func (tx *taxCalculator) visitTobacco(t *tobacco) {
    fmt.Printf("Price of tobacco item is = $%v \n", t.price*1.3)
 
}
 
func (tx *taxCalculator) visitNecessity(n *necessity) {
    fmt.Printf("Price of the necessity item is = $%v \n", n.price*1.13)
}
 
// concrete type for first nations tax calc behavior
type fntTaxCalculator struct {
    value float32
}
 
func (ftx *fntTaxCalculator) visitLiquor(l *liquor) {
    fmt.Printf("Price of first nation liquor item is = $%v \n", l.price*0.8)
}
 
func (ftx *fntTaxCalculator) visitTobacco(t *tobacco) {
    fmt.Printf("Price of first nation tobacco item is = $%v \n", t.price*0.5)
 
}
 
func (ftx *fntTaxCalculator) visitNecessity(n *necessity) {
    fmt.Printf("Price of the first nation necessity item is = $%v \n", n.price*1.03)
}
 
func main() {
    liquorItem := &liquor{price: 10.99}
    tobaccoItem := &tobacco{price: 10.99}
    necessityItem := &necessity{price: 10.99}
 
    taxCalculator := &taxCalculator{}
    liquorItem.accept(taxCalculator)
    tobaccoItem.accept(taxCalculator)
    necessityItem.accept(taxCalculator)
 
    fmt.Println()
    fntTaxCalculator := &fntTaxCalculator{}
    liquorItem.accept(fntTaxCalculator)
    tobaccoItem.accept(fntTaxCalculator)
    necessityItem.accept(fntTaxCalculator)
 
}
   
