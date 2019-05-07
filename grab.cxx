#include<iostream>
#include "grab.hxx"
#include<opencv4/opencv2/highgui.hpp>
#include <opencv2/imgproc/imgproc.hpp>

using namespace std;
using namespace cv;

Mat grabCutSegmentation(Mat  image,int x,int y,int width,int height){
  Mat bgModel,fgModel;
  Mat result;

  Rect rectangle(x,y,width,height);
 
    grabCut(image,    
            result,   
            rectangle,
            bgModel,fgModel, 
            1,        
            GC_INIT_WITH_RECT);


 compare(result,GC_PR_FGD,result,CMP_EQ);

 Mat foreGround(image.size(),CV_8UC3,Scalar(255,255,255));
 image.copyTo(foreGround,result);
  cv::rectangle(image, rectangle, Scalar(255,255,255),1);


  // namedWindow("Image");
//   imshow("Image",image);    
//   namedWindow("Segmented Image");
//   imshow("Segmented Image",foreGround);
 


// waitKey(0);
 
return foreGround;

}


int  foregroundDetection(int img_file,int x,int y,int width,int height){
string filename="image.jpg";

if(filename.empty()){
cout<<"File name is empty";
return 0;
}


cout<<"hi...im c++: "<<img_file<<endl;


Mat image=imread(filename,IMREAD_COLOR);

if(image.empty())
{
        cout<<"image is empty";
    return 0;
}

Mat result=grabCutSegmentation(image,x,y,width,height);
imshow("the real image",image);
waitKey(0);
imshow("result",result);
waitKey(0);


return 1;

}
