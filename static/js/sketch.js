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