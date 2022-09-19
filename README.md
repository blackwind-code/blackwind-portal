# blackwind-portal
Blackwind Service Portal for account registration &amp; VPN device management  
DGIST students can make blackwind account with @dgist.ac.kr email verification,  
and add devices to our VPN system through this portal  
  
# Reference
[Pocket Base Doc](https://pocketbase.io/docs/)

# Install
```bash
git clone https://github.com/blackwind-code/blackwind-portal.git
cd blackwind-portal/cmd/blackwind-portal
go build main.go

./main serve
```  
You need golang installed on your system  

# Add database schema
Register admin account on first startup (http://127.0.0.1:8090/_/)  
Admin page: Settings(on the left menu bar) -> Import collections -> Load from JSON file -> import "config/pb_schema.json"  
