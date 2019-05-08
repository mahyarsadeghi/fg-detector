#include<iostream>
#include "grab.hxx"
#include<opencv4/opencv2/highgui.hpp>
#include <opencv2/imgproc/imgproc.hpp>
#include<opencv2/opencv.hpp>
#include <curl/curl.h>
#include<string>
using namespace std;
using namespace cv;

size_t write_data(char *ptr, size_t size, size_t nmemb, void *userdata)
{
    vector<uchar> *stream = (vector<uchar>*)userdata;
    size_t count = size * nmemb;
    stream->insert(stream->end(), ptr, ptr + count);
    return count;
}

cv::Mat curlImg(const char *img_url, int timeout=10)
{
    vector<uchar> stream;
    CURL *curl = curl_easy_init();
     curl_easy_setopt(curl, CURLOPT_URL, img_url); 
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_data); 
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &stream); 
    curl_easy_setopt(curl, CURLOPT_TIMEOUT, timeout);
    CURLcode res = curl_easy_perform(curl); 
    curl_easy_cleanup(curl);
    return imdecode(stream, -1); 
}



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


int  foregroundDetection(const char  * img_url,int x,int y,int width,int height){
// const char * c= img_url.c_str();

// if(f.empty()){
// cout<<"File name is empty";
// return 0;
// }


cout<<"hi...im c++: "<<img_url<<endl;


Mat image=curlImg(img_url);

if(image.empty())
{
        cout<<"image is empty";
    return 1;
}

Mat result=grabCutSegmentation(image,x,y,width,height);
imshow("the real image",image);
waitKey(0);
imshow("result",result);
waitKey(0);


return 2;

}
