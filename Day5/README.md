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

## Info - Continuous Delivery (CD) Overview
<pre>
- is the matured level, where the QA certified builds are automatically deployed to live production environment
- in some companies, instead of making it live in production, they are deployed into pre-prod environment for further testing before it can be made live in production
</pre>
