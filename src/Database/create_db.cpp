#include <my_global.h>
#include <mysql.h>
#include "obtain_login.h"
void finish_with_error(MYSQL *con);
int main(int argc, char **argv)
{  
    MYSQL *con = mysql_init(NULL);

    if (con == NULL) 
    {
    fprintf(stderr, "%s\n", mysql_error(con));
    exit(1);
    }

    string User_name,password;
    User_name = Username();
    password = Password();
    
    int l = password.length();
    char pass_word[l+1];
    strcpy(pass_word,password.c_str());
    l = User_name.length();
    char username[l+1];
    strcpy(username,User_name.c_str());
    if (mysql_real_connect(con, "localhost", username, pass_word, NULL, 0, NULL, 0) == NULL) 
    {
        fprintf(stderr, "%s\n", mysql_error(con));
        mysql_close(con);
        exit(1);
    }  

    if (mysql_query(con, "show databases")) 
    {
        fprintf(stderr, "%s\n", mysql_error(con));
        printf("%ld products updated",(long) mysql_affected_rows(con));
        mysql_close(con);
        exit(1);
    }

     MYSQL_RES *result = mysql_store_result(con);
  
  if (result == NULL) 
  {
      finish_with_error(con);
  }

  int num_fields = mysql_num_fields(result);

  MYSQL_ROW row;
  
  while ((row = mysql_fetch_row(result))) 
  { 
      for(int i = 0; i < num_fields; i++) 
      { 
          printf("%s ", row[i] ? row[i] : "NULL"); 
      } 
          printf("\n"); 
  }
  
  mysql_free_result(result);

    mysql_close(con);
    exit(0);
}
void finish_with_error(MYSQL *con)
{
  fprintf(stderr, "%s\n", mysql_error(con));
  mysql_close(con);
  exit(1);        
}
