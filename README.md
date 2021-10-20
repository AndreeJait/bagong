# bagong

1. Generate template
   - add template base
   ```
   bagong template add --name project_name --value template_location
   ```
   - get template
   ```
   bagong template get --name project_name --folder template_folder --dest folderdestination
   ```
   - show all base
   ```
   bagong template list
   ```
   - delete base
   ```
   bagong template delete --name project_name
   ```
   - edit base
   ```
   bagong template edit --name project_name --value template_location
   ```
2. Generate resource
   - add resource base
   ```
   bagong resource add --name project_name --value github.com/user/example.git
   ```
   - get resource
   ```
   bagong resource get --name
   ```
   - show all base --name project_name --dest folderdestination
   ```
   bagong resource list
   ```
   - delete base
   ```
   bagong resource delete --name project_name
   ```
   - edit base
   ```
   bagong resource edit --name project_name --value github.com/user/example.git
   ```
