# Day 4

## Lab - Terraform modules
```
cd ~/terraform-may-2025
git pull
cd Day4/terraform/custom-terraform-module
terraform init
terraform apply --auto-approve
docker ps
```

Expected output
![image](https://github.com/user-attachments/assets/e9048a00-9068-439b-aaa4-9e1100f3d2d3)
![image](https://github.com/user-attachments/assets/78219f90-6d32-47e2-ba43-1e9bf434fdbb)
![image](https://github.com/user-attachments/assets/1723b206-6d69-4468-b1b5-f7ea09066c7f)
![image](https://github.com/user-attachments/assets/c7c52aa4-4434-4538-9b0f-0aaf322473ff)
![image](https://github.com/user-attachments/assets/b22e61b7-928d-407f-95f9-36bca1b3c377)
![image](https://github.com/user-attachments/assets/fb45d21b-314b-4861-a875-6300ce910425)
![image](https://github.com/user-attachments/assets/d33c3f20-4336-4d1b-b320-42948210f61c)
![image](https://github.com/user-attachments/assets/a4a56b7f-e77a-4358-95a0-e9383d7915fb)

Once you are done with this exercise, make sure you are disposing the resources
```
terraform destroy --auto-approve
```

Expected output
![image](https://github.com/user-attachments/assets/cc0206b7-886a-434f-aa19-ddcca6d02839)
![image](https://github.com/user-attachments/assets/67048628-2dc4-4c8d-8897-723a84232a2e)
![image](https://github.com/user-attachments/assets/409deae0-2e1a-4128-8e01-fc4757e406d2)
![image](https://github.com/user-attachments/assets/2c5b59c3-4833-4b1d-8b58-266d8138f487)


## Lab - Develop your first Custom Terraform Provider Plugin in Golang
You need to create a folder
```
mkdir -p /home/rps/go/bin
touch ~/.terraformrc
```
Paste the below code in your ~/.terraformrc file
```
provider_installation {
  dev_overrides {
    "registry.terraform.io/tektutor/file" = "/home/rps/go/bin",    
  }
  direct {}
}
```

Then you may proceed with the below instructions
```
cd ~/terraform-may-2025
git pull
cd Day4/custom-terraform-providers/file
go mod init github.com/tektutor/terraform-provider-file
go mod tidy
ls -l
go build
ls -l
go install
ls -l /home/rps/go/bin
```

Expected output
