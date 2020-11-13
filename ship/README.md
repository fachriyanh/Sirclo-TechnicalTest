#Object Oriented Programming GO
Go has no classes, no objects, no exceptions, and no templates like another languages. That's why GO has contrast to most object-oriented languages like C++, Java, C#, Scala, and even dynamic languages like Python and Ruby.

#Encapsulation
Go encapsulates things at the package level. Names that start with a lowercase letter are only visible within that package. You can hide anything in a private package and just expose specific types, interfaces, and factory functions. Vice versa, if you want to make the thing go public, give name that start with a uppercase

- Example : (Look at main.go)
    There I want to access two function. One to access the Distance, and want to access Code with another. But I cannot get the code because I don't have access to do that. Whereas I can access the Distance Info because that is public.

#Obejct Composition
Inheritance or subclassing was always a controversial issue. Modern languages and object-oriented thinking now favor composition over inheritance. Go takes it to heart and doesn't have any type hierarchy whatsoever. It allows you to share implementation details via composition.

-Example : (Look at kapal/kapal.go)
    There, I defined three struct like PerahuLayar, PerahuMotor, and KapalPesiar. Each struct has a SpekMesin component that refers to another struct that contains more detailed information like BahanBakar and KapasitasTanki. To give a value to that detail part, we can write (PerahuMotor).(SpekMesin).(BahanBakar). That is example how I put the BahanBakar Value for PerahuMotor

#Polymorphism
Polymorphism is the essence of object-oriented programming: the ability to treat objects of different types uniformly as long as they adhere to the same interface. Go interfaces provide this capability in a very direct and intuitive way.

-Example : (Look at kapal/kapal.go)
    There are three function in kapal.go. But functionally, they only give us two informations, Kode and JarakTemmpuh. They all have different calculation but I can write with the same function name. When you need that function for your object, Go will automatically choosed the right function to your object struct type as long as on the same interface.