//! trait parameter
trait Bark {
    fn bark(&self);
}

struct Dog;
impl Bark for Dog {
    fn bark(&self) {
        println!("woe!")
    }
}
struct Cat;
impl Bark for Cat {
    fn bark(&self) {
        println!("moew!")
    }
}

// method 1: use impl
fn make_bark(a: &impl Bark, b: &impl Bark) {
    a.bark();
    b.bark();
}

// method 2: use generic
//fn bark<T: Bark>(a: T, b: T) {
fn bark<T: Bark, U: Bark>(a: &T, b: &U) {
    a.bark();
    b.bark();
}

fn main() {
    let cat = Cat;
    let dog = Dog;
    make_bark(&cat, &dog);

    bark(&cat, &dog);
}
