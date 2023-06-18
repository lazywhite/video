#![allow(dead_code)]
// method1: 让编译器自动实现
// #[derive(Default)]
struct User {
    name: String,
    age: i32,
}

// method2: 手动实现Default Trait
/*
impl Default for User {
    fn default() -> Self {
        User {
            name: "template".to_owned(),
            age: 10,
        }
    }
}
*/

fn main() {
    let u1 = User {
        name: "bob".to_owned(),
        //..User::default()
    };
    println!("{}", u1.age);
}
