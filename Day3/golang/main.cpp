#include <iostream>
#include <string>
using namespace std;

//string sayHello( string );
string sayHello( string msg ) {
  return "Hello, " + msg + " !";
}

int main() {
	cout << sayHello("C++") << endl;
	return 0;
}

