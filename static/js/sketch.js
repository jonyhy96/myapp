var x=200;  
var y=200;  
var Vx=4;  
var Vy=3;  
var j=100;
var k=100;
var Vj=4;
var Vk=5;
function setup() {   
  createCanvas(400, 400);  
}   
  
function draw() {   
  background(222);  
  ellipse(x,y,20,20);//width和height是关键词，分别是Canvas的宽和高  
  ellipse(j,k,30,30);
	j+=Vj;
	k+=Vk;
	  if(j>width||j<0){  
        Vj*=-1;  
  }  
  if(k>height||k<0){  
        Vk*=-1;  
  }      
	x+=Vx;  
  y+=Vy;  
    if(x>width||x<0){  
        Vx*=-1;  
  }  
  if(y>height||y<0){  
        Vy*=-1;  
  }
}  
function setup() {
  createCanvas(400, 400);
  angleMode(DEGREES);
}

function draw() {
  background(0);
  translate(200, 200);
  rotate(-90);

  let hr = hour();
  let mn = minute();
  let sc = second();

  strokeWeight(8);
  stroke(255, 100, 150);
  noFill();
  let secondAngle = map(sc, 0, 60, 0, 360);
  //arc(0, 0, 300, 300, 0, secondAngle);

  stroke(150, 100, 255);
  let minuteAngle = map(mn, 0, 60, 0, 360);
  //arc(0, 0, 280, 280, 0, minuteAngle);

  stroke(150, 255, 100);
  let hourAngle = map(hr % 12, 0, 12, 0, 360);
  //arc(0, 0, 260, 260, 0, hourAngle);

  push();
  rotate(secondAngle);
  stroke(255, 100, 150);
  line(0, 0, 100, 0);
  pop();

  push();
  rotate(minuteAngle);
  stroke(150, 100, 255);
  line(0, 0, 75, 0);
  pop();

  push();
  rotate(hourAngle);
  stroke(150, 255, 100);
  line(0, 0, 50, 0);
  pop();

  stroke(255);
  point(0, 0);


  //  fill(255);
  //  noStroke();
  //  text(hr + ':' + mn + ':' + sc, 10, 200);


}