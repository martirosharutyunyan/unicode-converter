git add .
git status
read -p "Enter commit: " COMMIT
git commit -m "${COMMIT}"
git push