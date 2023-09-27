# Strategy_Pattern
Strategy Pattern Example
In this Go program, we demonstrate the Strategy Pattern, a behavioral design pattern. The Strategy Pattern defines a family of interchangeable 
algorithms and allows them to be used interchangeably within a context.

Overview
We implement different sorting algorithms: Bubble Sort, Merge Sort, and Insertion Sort, each as a separate "strategy" for sorting a collection 
of integers. These algorithms are encapsulated behind the Sorting interface, which defines a common Sort(arr []int) method. This design allows
us to choose and apply a sorting strategy at runtime without modifying the existing code.

How It Works
Interface: We define the Sorting interface, which declares a Sort method that all sorting strategies must implement. This interface acts as a 
contract that enforces the presence of the Sort method.

Concrete Strategies: Three concrete sorting strategies (bubbleSort, mergeSort, and insertionSort) implement the Sorting interface. Each strategy
encapsulates a specific sorting algorithm.

Factory Methods: Factory methods like NewBubbleSort(), NewMergeSort(), and NewInsertionSort() create instances of the respective sorting 
strategies, hiding the details of object creation.

Runtime Choice: In the main function, the user is prompted to choose a sorting algorithm by entering a number. Depending on the choice, the 
corresponding sorting strategy is instantiated and assigned to the sortMethod variable.

Polymorphism: We call the Sort() method on the sortMethod variable, which demonstrates polymorphism. The same method behaves differently 
based on the underlying sorting strategy.



Extensibility: Adding new sorting strategies is easy. You create a new sorting strategy that implements the Sorting interface, provide a factory method, and add a case in the switch statement. The existing code remains untouched.

Why It Matters
The Strategy Pattern promotes the separation of concerns and flexibility in software design. It allows us to decouple algorithms from their context and make them interchangeable. This pattern is beneficial when dealing with multiple algorithms that perform the same task but in different ways or when you want to provide users with choices for algorithm selection.

By encapsulating algorithms in separate strategies and using interfaces, we create a clean and extensible solution for sorting, demonstrating the power and elegance of the Strategy Pattern.
