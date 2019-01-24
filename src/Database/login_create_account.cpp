#include <my_global.h>
#include <mysql.h>
#include "obtain_login.h"

/**
 * 
 * Compile with ----g++ create_db.cpp -std=c99 `mysql_config --cflags --libs`------
 * 
**/
void finish_with_error(MYSQL *con);

int main(int argc, char **argv)
{  
    MYSQL *con = mysql_init(NULL);

    if (con == NULL) 
    {
    fprintf(stderr, "%s\n", mysql_error(con));
    exit(1);
    }

    printf("Enter DB details:\n");
    string User_name,password;
    User_name = Username();
    password = Password();
    
    //these are used to convert string to char array cuz the API needs it
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
        //mysql_close(con);
        exit(1);
    }

    MYSQL_RES *result = mysql_store_result(con);
  
    if (result == NULL) 
    {
        finish_with_error(con);
    }

    int num_fields = mysql_num_fields(result);

    MYSQL_ROW row;
    printf("Login to your account:\n");
    User_name = Username();
    
    
    int flag_login = 0;
    while ((row = mysql_fetch_row(result))) 
    { 
        for(int i = 0; i < num_fields; i++) 
        { 
            string s = row[i] ? row[i] : "NULL";
            if(s==User_name)
            {
                flag_login = 1;
                password = Password();
                printf("\nLogged in!\n");
            } 
        }
    }
    string bitch = "CREATE DATABASE " + User_name;
    l = bitch.length();
    char user_p[l+1];
    strcpy(user_p,bitch.c_str());

    if(flag_login==0)
    {
        printf("Account doesnt exist. Create new account by entering password\n");
        if (mysql_query(con, user_p)) 
        {
            printf("BITC");
            fprintf(stderr, "%s\n", mysql_error(con));
            printf("%ld products updated",(long) mysql_affected_rows(con));
            mysql_close(con);
            exit(1);
        }
        password = Password();
    }
    string user = "use " + User_name;
    
    l = user.length();
    char user_u[l+1];
    strcpy(user_u,user.c_str());

    if (mysql_query(con,user_u)) 
    {
        fprintf(stderr, "%s\n", mysql_error(con));
        printf("%ld products updated",(long) mysql_affected_rows(con));
        printf("Account created!\n");
        mysql_close(con);
        exit(1);
    }

    cout<<"WTF";
    if (mysql_query(con,"show tables")) 
    {
        fprintf(stderr, "%s\n", mysql_error(con));
        printf("%ld products updated",(long) mysql_affected_rows(con));
        printf("Account created!\n");
        mysql_close(con);
        exit(1);
    }

while ((row = mysql_fetch_row(result))) 
    { 
        for(int i = 0; i < num_fields; i++) 
        { 
            string s = row[i] ? row[i] : "NULL";
            if(s==User_name)
            {
                flag_login = 1;
                password = Password();
                printf("\nLogged in!\n");
            } 
        }
    }
    mysql_free_result(result);
    mysql_close(con);
    exit(0);

    return 0;
}

void finish_with_error(MYSQL *con)
{
fprintf(stderr, "%s\n", mysql_error(con));
mysql_close(con);
exit(1);        
}
