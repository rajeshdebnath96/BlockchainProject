#nvm codes
nvm install 8.12.0
nvm use node



./startFabric.sh
node enrollAdmin.js
node registerUser.js
node query.js




#Error: listen EADDRINUSE :::3001
#pid=`sudo lsof -n -i :3001|grep LISTEN|awk '{print $2}'`
#echo $pid hjkhh
#kill -9 `sudo lsof -n -i :3001|grep LISTEN|awk '{print $2}'`


#git commands
# git init
# git add <> -A for all folder_name/
# git commit -m "first commit"
# git remote add origin https://github.com/rajeshdebnath96/BlockchainProject.git
# git remote set-url origin https://github.com/rajeshdebnath96/BlockchainProject.git
# git push -u origin master
# git push -u --force origin master
# git pull origin master
# git rm --cache -f -r my_codes/

# storing git username & password
# git config credential.helper store 
# git pull

