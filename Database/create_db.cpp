#include <my_global.h>
#include <mysql.h>
#include "obtain_login.h"

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

    if (mysql_query(con, "CREATE DATABASE testdb")) 
    {
        fprintf(stderr, "%s\n", mysql_error(con));
        mysql_close(con);
        exit(1);
    }

    mysql_close(con);
    exit(0);
}