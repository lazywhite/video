#[derive(Debug)]
struct Node {
    value: i32,
    left: Option<Box<Node>>,
    right: Option<Box<Node>>,
}
fn main() {
    let n2 = Node {
        value: 2,
        left: None,
        right: None,
    };
    let root = Node {
        value: 1,
        left: Some(Box::new(n2)),
        right: None,
    };
    println!("{:?}", root);
}
