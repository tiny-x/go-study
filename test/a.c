#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/types.h>
#include <pwd.h>

int main()
  {
    struct passwd *pw;

    if( ( pw = getpwuid( getuid() ) ) == NULL ) {
       fprintf( stderr,
          "getpwuid: no password entry\n" );
       exit( EXIT_FAILURE );
    }
    printf( "login name  %s\n", pw->pw_name );
    printf( "user ID     %d\n", pw->pw_uid );
    printf( "group ID    %d\n", pw->pw_gid );
    printf( "home dir    %s\n", pw->pw_dir );
    printf( "login shell %s\n", pw->pw_shell );
    exit( EXIT_SUCCESS );
  }