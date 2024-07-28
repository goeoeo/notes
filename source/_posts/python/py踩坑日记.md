# mac pip安装psycopg2报错
 Error: pg_config executable not found.   

1. 安装
    ```
    # 记得换源 
    brew install postgresql
    ```
2. vim ~/.zshrc
    ```
    export PATH="/opt/homebrew/opt/apr-util/bin:$PATH"
    export LDFLAGS="-L/opt/homebrew/opt/apr-util/lib"
    export CPPFLAGS="-I/opt/homebrew/opt/apr-util/include"
    export PKG_CONFIG_PATH="/opt/homebrew/opt/apr-util/lib/pkgconfig"
    
    ```
3. pip install psycopg2==2.9.9
