/*
 *  NAME: ScreenSaver Mouse Jiggler
 *  DATE: 2016-10-20
 *  DESC: Arduino based mouse emulator, preventing computer screen-saver from
 *      kicking in and locking desktop (eg. during forensic investigation).
 *  AUTHOR: nshadov
 *  VERSION: 1.0
 */
 

#include <Mouse.h>

void setup()
{
  Mouse.begin();
}


void loop()
{
  delay(1000);

  while(true) { 
    Mouse.move(4,0,0);
    delay(100);
    Mouse.move(-8,0,0);
    delay(100);
    Mouse.move(4,0,0);
    
    delay(5000);
  }
}
