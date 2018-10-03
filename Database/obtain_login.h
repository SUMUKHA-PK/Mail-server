#include<bits/stdc++.h>
#include "echo.h"
using namespace std;

string Username()
{
    string username;
    cout<<"Enter Username: ";
    cin>>username;
    return username;
}
string Password()
{
    string password;
    cout<<"Enter password: ";
    echo_off();
    cin>>password;
    echo_on();
    return password;
}