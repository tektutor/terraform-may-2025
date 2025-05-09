# Day 5

## Info - DevOps Overview
<pre>
- automating end-2-end 
- convert dev testing, QA testing, prod test into code so that the code can be pused into version control
- when CI Build Servers detects code commit, it can grab the latest code and trigger automated builds that includes automated test cases
- idea is to catch defects early
</pre>  

## Info - Continuous Integration (CI) Overview
<pre>
- Fail-fast engineering process
- we need self-managing professional team
- as soon as some code is ready with automated test case, code code will merged with dev branch several times a time
- the intent of automated CI builds is to catch the defects early, if possible during the development cycle itself
- if the build fails after merging your code
  - due some test case failures, they are considered a good
- with the help of CI build with automated test cases, we can release the product frequently with confidence
- end-to-end build process should be automated with Jenkins, TeamCity, bamboo, etc.,
- usually TDD(Test Driven Development) approach is used to automated test and develop features and fix bugs
</pre>

## Info - Continuous Deployment (CD) Overview
<pre>
- the dev certified binaries, will be automatically deployed onto QA environments for further automated tests
- end-2-end functionality test, stress/load test, component test, performance test everything will be automated using BDD
</pre>

## Jenkins Overview
<pre>
- very first Continuous Integration Build server 
- Josuke Kawaguchi is the person who developed Jenkins
- initially it was called Hudson, it was developed in Java while he was working for Sun Microsystems
- many like minded colleagues started contributed and started using Hudson for automating Build within Sun Microsystems
- meanwhile, Oracle acquired Sun Microsystems, and Oracle had some proposals to make Hudson a commercial product
- the team that was developing Hudson didn't like the idea of Oracle making the producat a commercial one, hence most of the team members including Josuke Kawaguchi had quit Oracle and Josuke founded a company called Cloudbees forking the Hudson branch into a new branch called Jenkins
- The baseline code for Jenkins comes from Hudson, even today in some of the logs you could see the Hudson word appearing
- the team stayed back in Oracle continues to develop/maintain the Hudson product, but Jenkins is the most popular CI Build server even today
- inspired by Jenkins many other CI Builds servers started coming up which includes TeamCity, Bamboo, Circle City, TFS, etc.,
- Jenkins is opensource
- the enterprise variant of Jenkins is called Cloudbees
- while Cloudbees is a commercial enterprise product that requires license, but functionally it resembles Jenkins
  
</pre>

## Info - Continuous Delivery (CD) Overview
<pre>
- is the matured level, where the QA certified builds are automatically deployed to live production environment
- in some companies, instead of making it live in production, they are deployed into pre-prod environment for further testing before it can be made live in production
</pre>

## Lab - Launching Jenkins from your Linux Terminial
```
cd ~/Downloads
ls -l jenkins.war

java -jar ./jenkins.war
```

Expected output
![image](https://github.com/user-attachments/assets/8a00b093-0881-4fc6-8ad1-04a4077c4040)
![image](https://github.com/user-attachments/assets/e29ce98b-edbd-498d-b0c0-219c464e84da)

You can the Jenkins Dashboard from RPS Cloud Machine (chrome web browser )
<pre>
http://localhost:8080  
</pre>
![image](https://github.com/user-attachments/assets/0d116ec1-0c42-42a7-b0c6-2db69d0dd92b)
![image](https://github.com/user-attachments/assets/28338c55-8c5f-4a87-9879-3f0522cfe4e2)
                                                                                                      
