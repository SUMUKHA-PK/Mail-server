#include<bits/stdc++.h>
#include <mysql_connection.h>
#include <cppconn/driver.h>
#include <cppconn/exception.h>
#include <cppconn/resultset.h>
#include <cppconn/statement.h>
#include "obtain_login.h"

using namespace std;

int main()
{
  try {
  sql::Driver *driver;
  sql::Connection *con;
  sql::Statement *stmt;
  sql::ResultSet *res;

  string User_name,password;
  User_name = Username();
  password = Password();

  driver = get_driver_instance();                                       //Create a connection to mysql server

  con = driver->connect("localhost", User_name, password);
  con->setSchema("SCA");                                                /* Connect to the MySQL test database */

  stmt = con->createStatement();
  res = stmt->executeQuery("CREATE DATABASE impk1");        // replace with your statement
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