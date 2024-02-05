# Password Manager
This is a CLI tool I am writing to:
1. Gain experience with Cobra and Viper for my job.
2. Continue learning Go.
3. Generate and manage passwords as I am terrible about this.

# What is Available?
- Ability to generate passwords for specific applications (not persistent yet)
- Example: `password-go create [APPLICATION] -s=true -l=15` where s = special and l = length

# What is on the Horizon?
- The ability to actually persist passwords between interactions with the tool
- Being able to view all passwords or a single password
- Login / credentials before creation or viewing of passwords
- Some form of encryption so passwords are not in plaintext
- Updating and deleting passwords