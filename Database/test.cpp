#include<bits/stdc++.h>
#include <mysql_connection.h>
#include <cppconn/driver.h>
#include <cppconn/exception.h>
#include <cppconn/resultset.h>
#include <cppconn/statement.h>
#include <termios.h>
void echo_off(void);
void echo_on(void);
static struct termios stored_settings;
using namespace std;
int main()
{
  try {
  sql::Driver *driver;
  sql::Connection *con;
  sql::Statement *stmt;
  sql::ResultSet *res;

  string Username,password;
  cout<<"Enter Username: ";
  cin>>Username;
  cout<<"Enter password: ";
  echo_off();
  cin>>password;
  echo_on();
 
  driver = get_driver_instance();                                       //Create a connection to mysql server

  con = driver->connect("localhost", Username, password);
  con->setSchema("SCA");                                                /* Connect to the MySQL test database */

  stmt = con->createStatement();
  res = stmt->executeQuery("SELECT 'Hello World!' AS _message");        // replace with your statement
  while (res->next()) {
    cout << "\n... MySQL replies: ";
    cout << res->getString("_message") << endl;                         /* Access column data by alias or column name */
    cout << "... MySQL says it again: ";
    cout << res->getString(1) << endl;                                  /* Access column fata by numeric offset, 1 is the first column */
  }
  delete res;
  delete stmt;
  delete con;

} catch (sql::SQLException &e) {
  cout << "# ERR: SQLException in " << __FILE__;
  cout << "(" << __FUNCTION__ << ") on line "<< __LINE__ << endl;
  cout << "# ERR: " << e.what();
  cout << " (MySQL error code: " << e.getErrorCode();
  cout << ", SQLState: " << e.getSQLState() << " )" << endl;
}

  return 0;
}

void echo_off(void)                                                    //Switches off Echo for password
{
  struct termios new_settings;
  tcgetattr(0,&stored_settings);
  new_settings = stored_settings;
  new_settings.c_lflag &= (~ECHO);
  tcsetattr(0,TCSANOW,&new_settings);
  return;
} 

void echo_on(void)                                                    //Switch echo bak on in the terminal
{
  tcsetattr(0,TCSANOW,&stored_settings);
  return;
}