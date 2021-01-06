<div align="center"> 
<!-- <img align="center" alt="algorithm-visualizer" src="Logo1.png" height='150' width='150'> -->
<h1 align="center"> Library</h1>

A web application, that lets you to share your regular reads and your books with your friends and keeps you updated with your friends activiteis as well.

<!-- ![Issues](https://img.shields.io/github/issues/servermonk/Algorithm-Visualizer)
![PRs](https://img.shields.io/github/issues-pr/servermonk/Algorithm-Visualizer) -->
![Maintenance](https://img.shields.io/maintenance/yes/2020)
![star](https://img.shields.io/github/stars/anubhavitis/Library?style=social)
![fork](https://img.shields.io/github/forks/anubhavitis/Library?style=social)


</div>
<hr>

## Getting Started

<!-- > Click [here](https://vizalgo.netlify.app/) to get directed to the hosted page. -->
> Application is under development phase, Coming Soon!

<!-- #### GitHub Repository Structure

| S.No. | Branch Name                                                              | Purpose                       |
| ----- | ------------------------------------------------------------------------ | ----------------------------- |
| 1.    | [master](https://github.com/servermonk/Algorithm-Visualizer/tree/master) | contains the main code        |
| 2.    | [beta](https://github.com/servermonk/Algorithm-Visualizer/tree/beta)     | contains all development code |

> *Note*: We're not accepting any changes in the `master` branch, make PRs in the `beta` branch only. -->

## Technology Stack Used
<img src="https://img.shields.io/badge/html5%20-%23E34F26.svg?&style=for-the-badge&logo=html5&logoColor=white"/> <img src="https://img.shields.io/badge/css3%20-%231572B6.svg?&style=for-the-badge&logo=css3&logoColor=white"/> <img src="https://img.shields.io/badge/javascript%20-%23323330.svg?&style=for-the-badge&logo=javascript&logoColor=%23F7DF1E"/> <img src="https://img.shields.io/badge/GO%20-%2343853D.svg?&style=for-the-badge&logo=GOLANG&logoColor=white"/> <img src="https://img.shields.io/badge/github%20-%23121011.svg?&style=for-the-badge&logo=github&logoColor=white"/>

- <strong> FrontEnd Design: </strong> HTML & CSS
- <strong> Algorithms and animations: </strong> JavaScript
- <strong> Server hosting: </strong> GO Web Framework
- <strong> Web hosting: </strong> Heroku

#### Version of API :  V1.0
    
# API Documentation
 All responses come in standard JSON. All requests must include a `content-type` of `application/json` and the body must be valid JSON.


### Response Codes
```
200: Success
201: Created
400: Bad request
401: Unauthorized
404: Cannot be found
405: Method not allowed
50x: Server Error
```
### Error and Success Message Example

```json
  {
    "error":"" 
  }
  
  {
      "success":"",
      "data": ""  
  }
```

## Custom SignUp

**You send:**  You send the details required to signup.

**You get:** An `Error-Message` or a `Success-Message` depending on the status of the account created

**Endpoint:** 
     /signup

**Authorization Token:** Not required

**Request:**
`POST HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{   
    "name": "abc",
    "username": "user",
    "email": "foo",
    "password": "1234567",
    "cpassword": "1234567"
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
   "success": true,
   "Token" : "hgsvcdvbhjcu76tdghev3bneidf87ydtegwvb3e"
}
```

## Google signup

**You send:**  Your credentials and authorization to the app.

**You get:** An `API-Token` and a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /google/signup

**Authorization Token:** Not required

**Request:**  sent to google oauth.

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
   "success":true,
   "token": "e3b0iuytfgvbjkio876yghhjio9098765rtfgvb",
}
```

## Twitter signup

**You send:**  Your credentials and authorization to the app.

**You get:** An `API-Token` and a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /twitter/signup

**Authorization Token:** Not required

**Request:**  sent to twitter oauth.

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
   "success":true,
   "token": "e3b0iuytfgvbjkio876yghhjio9098765rtfgvb",
}
```
<!-- ## Project Setup
- Fork and clone the Repo by typing the following commands in the terminal 
```
$ git clone https://github.com/your-username/Algorithm-Visualizer.git
$ cd Algorithm Viusalizer
```
- Change Branch using:
```
$ git checkout beta
```
- To open the site you can either use [Live Server](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer) of VS-Code or similar tools, or you need to install Node.js 
    <details>
    To open site on Localhost:
    - Install node dependencies using:

    ```
    $ npm install
    ```

    - To start the server, type:
    ```
    $ node server
    ```
   
    - Then on your browser type http://localhost:3000/
  </details>
- Make changes to the code and save your changes
- Commit your changes using:
```
$ git commit -m "add any comment"
```
- Push the changes to the forked repository
- Navigate to the original repository and make a pull request -->


<!-- ## Project administrators ✨ -->

<!-- Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)): -->


<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<!-- <table>
  <tbody>
    <tr>
      <td align="center">
        <a href="https://github.com/anubhavitis"><img alt="Anubhav" src="https://avatars1.githubusercontent.com/u/26124625?s=400&u=c411643ffe3db941107eca578ada396c5f8dfa3a&v=4" width="100px;"><br><sub><b>Anubhav Singhal</b></sub></a><br>
      </td>
      <td align="center">
        <a href="https://github.com/vashish1">
          <img alt="Yashi" src="https://user-images.githubusercontent.com/26124625/103191314-8d06f780-48fa-11eb-8553-56d52a6f8faf.jpeg" width="100px;"> <br>
          <sub><b>Yashi Gupta</b></sub></a>  <br>
        </a>
      </td>
    </tr>
  </tbody>
</table> -->

<!-- ## Beloved Contributors ✨

Thanks goes to these wonderful people:


[![](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/images/0)](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/links/0)[![](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/images/1)](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/links/1)[![](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/images/2)](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/links/2)[![](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/images/3)](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/links/3)[![](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/images/4)](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/links/4)[![](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/images/5)](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/links/5)[![](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/images/6)](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/links/6)[![](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/images/7)](https://sourcerer.io/fame/servermonk/servermonk/Algorithm-Visualizer/links/7) -->


## **Thank You**
> Made with Love ❤️️  &  Passion 🙏.
