#include <termios.h>
static struct termios stored_settings;

void echo_off(void)                                                    //Switches off Echo for password
{
  struct termios new_settings;
  tcgetattr(0,&stored_settings);
  new_settings = stored_settings;
  new_settings.c_lflag &= (~ECHO);
  tcsetattr(0,TCSANOW,&new_settings);
  return;
} 

void echo_on(void)                                                    //Switch echo back on in the terminal
{
  tcsetattr(0,TCSANOW,&stored_settings);
  return;
}