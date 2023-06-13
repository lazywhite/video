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

fn get_animal(num: i32) -> Box<dyn Bark> {
    if num % 2 == 0 {
        return Box::new(Dog);
    }
    return Box::new(Cat);
}

// fn wrong(num: i32) -> impl Bark {
//     if num % 2 == 0 {
//         return Dog;
//     }
//     return Cat;
// }

fn main() {
    let a = get_animal(1);
    a.bark();
    let b = get_animal(2);
    b.bark();
}
