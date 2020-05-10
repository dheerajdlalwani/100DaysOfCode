//  Variables
// var, let & const

//var
var a;
a = 10;
var b = 20;
var c = a * b;
console.log("a = " + a );
console.log("b = " + b );
console.log("c = a * b = " + c );

//let
let d;
d = 10;
let e = 20;
let f = d * e;
console.log("d = " + d );
console.log("e = " + e );
console.log("f = d * e = " + f );
    //re-assinging 
    console.log("Re-assining the values of variables:")
    d = 20;
    e = 5;
    f = d * e;
    console.log("d = " + d );
    console.log("e = " + e );
    console.log("f = d * e = " + f );


//const
const g = 10;
//g = 20;     //uncomment to see the error
console.log("g = " + g );
    //Error: 
    //Uncaught TypeError: Assignment to constant variable.

//---------------------------------------------------------------------------------

// Data-Types
// string, number, boolean, null, undefined

//string
const username = 'HackMeEasy101';
console.log("username = " + username);
console.log("Type of username = " + typeof username);

//number
const age = 18;
console.log("age = " + age);
console.log("Type of age = " + typeof age);
const score = 95.5; //There is nothing like float
console.log("score = " + score);
console.log("Type of score = " + typeof score);

//boolean
const pass = true;
console.log("Is the student passed? Answer: " + pass);
console.log("Type of pass = " + typeof pass);

//null
const emptyBox = null;
console.log("emptyBox = " + emptyBox);
console.log("Type of emptyBox = " + typeof emptyBox);

//undefined
const notYetDecided = undefined;
console.log("age = " + age);
console.log("Type of notYetDecided = " + typeof notYetDecided);
let notYetDecided2;
console.log("notYetDecided2 = " + notYetDecided2);
console.log("Type of notYetDecided2 = " + typeof notYetDecided2);

//string methods
const example = "This is amazing!";
console.log("example = " + example);
console.log(`Length of '${example}' is ${example.length}`);
console.log(`'${example}' in all CAPS = '${example.toUpperCase()}'`);
console.log(`Substring or slice of '${example}' = '${example.substring(0, 7)}'`);
console.log(`Array of each character of '${example}' = '${example.split('')}'`);
console.log(`Array of each word of '${example}',  seperated by space = '${example.split(' ')}'`);

//arrays
const mortalEnemies = ['MU', 'TSEC', 'Mumbai', 'Maharashtra', 'Local Trains', 'Crowd'];
console.log(`My list of mortal enemies = [${mortalEnemies}]`);
console.log(`3rd enemy in my list = ${mortalEnemies[2]}`);
mortalEnemies.push('Snakes');
console.log(`My new list of mortal enemies = [${mortalEnemies}]`);
mortalEnemies.unshift('Non-Vegetarians');
console.log(`My new list of mortal enemies = [${mortalEnemies}]`);
mortalEnemies.pop();
console.log(`My new list of mortal enemies = [${mortalEnemies}]`);
console.log(`Index of 'Mumbai' in my list = ${mortalEnemies.indexOf('Mumbai')}`);

// Object Literals
const coder = {
    firstName: 'Robert',
    lastName: 'Zane',
    age: 57,
    weaponOfChoice: 'Python',
    arsenal: ['Python', 'HTML', 'CSS', 'JS', 'C++', 'Java'],
    projects: {
        Android: 'Resource Manager',
        WebDev: 'Customer Management System',
        DataAnalysis: 'COVID-19 Prediction System'
    }
}
console.log(`The properties of my coder are = `);
console.log(`First Name = ${coder.firstName}`);
console.log(`Last Name = ${coder.lastName}`);
console.log(`Age = ${coder.age}`);
console.log(`Weapon of choice = ${coder.weaponOfChoice}`);
console.log(`Arsenal Includes = [${coder.arsenal}]`);
console.log(`Projects are = [${coder.projects.Android}, ${coder.projects.DataAnalysis} & ${coder.projects.WebDev}]`);
coder.email = 'zanerobzane@gmail.com';
console.log(`Email: ${coder.email}`);

// JSON
const coderJSON = JSON.stringify(coder);
console.log(coderJSON);

// LOOPS
for(let z = 0; z <= 7; z++)
{
    console.log(`Hit ${z}`);
}
let y = 0;
while(y<=7)
{
    console.log(`Big Hit ${y}`);
    y = y + 1;
}
 
//Class
class CodeMaster
{
    constructor(fName, lName, dob, lang)
    {
        this.fName = fName;
        this.lName = lName
        this.dob = new Date(dob);
        this.lang = lang;
    }
    getBirthYear()
    {
        return this.dob.getFullYear;
    }
    getFullName() 
    {
        return `${this.fName} ${this.lName}`;
    }
    getLanguage()
    {
        return this.lang;
    }
}

//Instantiate Objects
const master1 = new CodeMaster('Mike', 'Ross', '03-22-2006', 'Python');
const master2 = new CodeMaster('Rachel', 'Zane', '03-22-2008', 'JavaScript');
console.log(master1);
console.log(master2);
console.log(master1.getLanguage());
console.log(master2.getLanguage());
