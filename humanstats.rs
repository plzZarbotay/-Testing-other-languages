use std::io;
 
fn main()  -> io::Result<()>{

    let mut input = String::new();
     let mut inputage =String::new();
     let mut project = String::new();
      let mut ling = String::new();
    println!("Введите свое имя: ");
    io::stdin().read_line(&mut input)?;
    println!("Введите ваш возраст: ");
    io::stdin().read_line(&mut inputage);
    println!("Введите ссылку на ваш код");
    io::stdin().read_line(&mut project);
  println!("Введите Ваши данные:");
    io::stdin().read_line(&mut ling);
    println!("Ваш имя: {}", input);
    println!("Ваш возраст: {}", inputage);
     println!("Ссылка на ваш код: {}", project);
     println!("Ссылка на ваш код: {}", ling);
    Ok(())
}

