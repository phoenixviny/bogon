###git tourial 
#####	initialliazie
-	$git clone username@host:/path/to/repository   /*clone a exist repository*/
-	$mkdir repo && cd repo  $git init              /*init a repository named repo*/
-	$git init <dir name>						   /*transfer a normal dir to git repository*/
	in order to manger the local repository,git have three different tree, there is “working dir”、"index or stage"、
"HEAD for last commit of localstorage"
-	$git add <a file name>							/*add a file to index*/
-	$git add .										/*add all unstaged file to index*/
-	$git commit -m "meaningful commit message"		/*commit changes to HEAD*/
-	$git remote add origin <server>					/*touch a local repo to remote repo*/
-	$git push origin master							/*push HEAD to remote repo*/
	to show the status of working dir and index 's changes and these files has being tracking,which hasn't
-	$git status 									/*show the current working dir's status,seeing which not being watched*/
-	$touch .gitignore								/*initialliazie a ignore file that git identel*/
-	$git log										/*using default formate show a full git history*/
-	$git log -n 3									/*only show last 3 commits dir working logs*/
-	$git log --pretty=oneline 						/*compoise log file to a line,show working commit logs*/
-	$git log -p										/*show every commit and  the changes of it,most clear show logs of git repo*/
-	$git log --author="a name or a regx rule"		/*using var search the commit log,that var string can be a name or a regx*/
-	$git log --grep="a string or a regx rule"		/*search commit logs that fits a string or a regx rule*/
-	$git log --graph --decorate --oneline			/*show logs using graphic img*/
	therefor has a mehod to checkout a file and old commit and checkout a branch from current working dir
-	$git checkout master							/*backoff to master branch*/
-	$git checkout <commit> <filename>				/*checkout a file from the commit,drop the edit in the work dir*/
-	$git checkout <commit>							/*drop all files'changes,and clone the context from commit*/
	fallback the err editing,revert:fallback a solo commit,don't remove the commit below it,first revert don't changing working history
for these commit that already pubilc a share repo
-	$git revert <commit>							/*undo a changes from comit with current branch*/


























