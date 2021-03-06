#!/usr/bin/env python
# coding: utf-8

# In[1]:


VERSION = "0.0.6"
version = VERSION
__version__ = VERSION


# In[2]:


import os, sys, json, resource
import pandas as pd
import numpy as np
from random import shuffle
import cv2


# In[3]:


import tensorflow as tf
from tensorflow.keras.preprocessing.image import ImageDataGenerator


# In[4]:


from imgaug import augmenters as iaa
import imgaug as ia
ia.seed(1)


# In[5]:


def getmem():
    print('Memory usage         : % 2.2f MB' % round(
        resource.getrusage(resource.RUSAGE_SELF).ru_maxrss/1024.0/1024.0,1)
    )


# In[6]:


def get_id_from_file_path(file_path, suffix_ = False):
    if suffix_: 
        return file_path.split(os.path.sep)[-2] + '/' + file_path.split(os.path.sep)[-1].replace('.png', '')
    else:
        return file_path.split(os.path.sep)[-1].replace('.png', '')


# In[7]:


def chunker(seq, size):
    return (seq[pos:pos + size] for pos in range(0, len(seq), size))


# In[8]:


def MultiClass(seq_proc):
    def wrapper(list_files, 
                id_label_map, 
                batch_size,
                depth,
                suffix = False,
                augment=False,
                shuf = False):
        X = np.full(batch_size, 0).astype(float)
        while True:
            if shuf: shuffle(list_files)
            for batch in chunker(list_files, batch_size):
                for i, x in enumerate(batch):                    
                    image= cv2.imread(x)
                    yy = id_label_map[get_id_from_file_path(x, suffix_ = suffix)]
                    if i == 0:
                        X = np.expand_dims(np.array(image, dtype= 'uint8'), axis= 0)
                        Y = np.expand_dims(np.array(yy, dtype= 'uint8'), axis= 0)
                    else:
                        image= np.expand_dims(np.array(image, dtype= 'uint8'), axis= 0)
                        yy = np.expand_dims(np.array(yy, dtype= 'uint8'), axis= 0)
                        X = np.append(X, image, axis= 0)
                        Y = np.append(Y, yy, axis= 0)
                if augment:
                    X = seq_proc(list_files, 
                                id_label_map, 
                                batch_size, 
                                augment).augment_images(X)
                X = X/255.0 
                
                yield np.array(X), tf.one_hot(np.array(Y), depth = depth)
    return wrapper


# In[9]:


def LinRegr(seq_proc):
    def wrapper(list_files, 
                id_label_map, 
                batch_size,
                depth,
                augment=False,
                shuf = False):
        X = np.full(batch_size, 0).astype(float)
        while True:
            if shuf: shuffle(list_files)
            for batch in chunker(list_files, batch_size):
                for i, x in enumerate(batch):                    
                    image= cv2.imread(x)
                    yy = id_label_map[get_id_from_file_path(x)]
                    if i == 0:
                        X = np.expand_dims(np.array(image, dtype= 'uint8'), axis= 0)
                        Y = np.expand_dims(np.array(yy, dtype= 'uint8'), axis= 0)
                    else:
                        image= np.expand_dims(np.array(image, dtype= 'uint8'), axis= 0)
                        yy = np.expand_dims(np.array(yy, dtype= 'uint8'), axis= 0)
                        X = np.append(X, image, axis= 0)
                        Y = np.append(Y, yy, axis= 0)
                if augment:
                    X = seq_proc(list_files, 
                                id_label_map, 
                                batch_size, 
                                augment).augment_images(X)
                X = X/255.0 
                
                yield np.array(X), np.array(Y)
    return wrapper


# In[10]:


#light augmenter, linear regression
@LinRegr
def LightImgAugDataGeneratorLR(list_files, 
        id_label_map, 
        batch_size, 
        augment=False,
        shuf = False):
    
    return iaa.Sequential([
    iaa.Fliplr(0.5), # horizontal flips
    iaa.Crop(percent=(0, 0.1)), # random crops
    # Small gaussian blur with random sigma between 0 and 0.5.
    # But we only blur about 50% of all images.
    iaa.Sometimes(
        0.5,
        iaa.GaussianBlur(sigma=(0, 0.5))
    ),
    # Strengthen or weaken the contrast in each image.
    iaa.LinearContrast((0.75, 1.5)),
    # Add gaussian noise.
    # For 50% of all images, we sample the noise once per pixel.
    # For the other 50% of all images, we sample the noise per pixel AND
    # channel. This can change the color (not only brightness) of the
    # pixels.
    iaa.AdditiveGaussianNoise(loc=0, scale=(0.0, 0.05*255), per_channel=0.5),
    # Make some images brighter and some darker.
    # In 20% of all cases, we sample the multiplier once per channel,
    # which can end up changing the color of the images.
    iaa.Multiply((0.8, 1.2), per_channel=0.2),
    # Apply affine transformations to each image.
    # Scale/zoom them, translate/move them, rotate them and shear them.
    iaa.Affine(
        scale={"x": (0.8, 1.2), "y": (0.8, 1.2)},
        translate_percent={"x": (-0.2, 0.2), "y": (-0.2, 0.2)},
        rotate=(-25, 25),
        shear=(-8, 8)
    )
    ], random_order=True) # apply augmenters in random order


# In[11]:


#light augmenter, multiclass
@MultiClass
def LightImgAugDataGeneratorMC(list_files, 
        id_label_map, 
        batch_size, 
        augment=False,
        shuf = False):
    
    return iaa.Sequential([
    iaa.Fliplr(0.5), # horizontal flips
    iaa.Crop(percent=(0, 0.1)), # random crops
    # Small gaussian blur with random sigma between 0 and 0.5.
    # But we only blur about 50% of all images.
    iaa.Sometimes(
        0.5,
        iaa.GaussianBlur(sigma=(0, 0.5))
    ),
    # Strengthen or weaken the contrast in each image.
    iaa.LinearContrast((0.75, 1.5)),
    # Add gaussian noise.
    # For 50% of all images, we sample the noise once per pixel.
    # For the other 50% of all images, we sample the noise per pixel AND
    # channel. This can change the color (not only brightness) of the
    # pixels.
    iaa.AdditiveGaussianNoise(loc=0, scale=(0.0, 0.05*255), per_channel=0.5),
    # Make some images brighter and some darker.
    # In 20% of all cases, we sample the multiplier once per channel,
    # which can end up changing the color of the images.
    iaa.Multiply((0.8, 1.2), per_channel=0.2),
    # Apply affine transformations to each image.
    # Scale/zoom them, translate/move them, rotate them and shear them.
    iaa.Affine(
        scale={"x": (0.8, 1.2), "y": (0.8, 1.2)},
        translate_percent={"x": (-0.2, 0.2), "y": (-0.2, 0.2)},
        rotate=(-25, 25),
        shear=(-8, 8)
    )
    ], random_order=True) # apply augmenters in random order


# In[12]:


#deep augmenter
@LinRegr
def DeepImgAugDataGeneratorLR(list_files, 
        id_label_map, 
        batch_size, 
        augment=False,
        shuf = False):
    
    sometimes = lambda aug: iaa.Sometimes(0.5, aug)
    return iaa.Sequential(
    [
        #
        # Apply the following augmenters to most images.
        #
        iaa.Fliplr(0.5), # horizontally flip 50% of all images
        iaa.Flipud(0.2), # vertically flip 20% of all images

        # crop some of the images by 0-10% of their height/width
        sometimes(iaa.Crop(percent=(0, 0.1))),

        # Apply affine transformations to some of the images
        # - scale to 80-120% of image height/width (each axis independently)
        # - translate by -20 to +20 relative to height/width (per axis)
        # - rotate by -45 to +45 degrees
        # - shear by -16 to +16 degrees
        # - order: use nearest neighbour or bilinear interpolation (fast)
        # - mode: use any available mode to fill newly created pixels
        #         see API or scikit-image for which modes are available
        # - cval: if the mode is constant, then use a random brightness
        #         for the newly created pixels (e.g. sometimes black,
        #         sometimes white)
        sometimes(iaa.Affine(
            scale={"x": (0.8, 1.2), "y": (0.8, 1.2)},
            translate_percent={"x": (-0.2, 0.2), "y": (-0.2, 0.2)},
            rotate=(-45, 45),
            shear=(-16, 16),
            order=[0, 1],
            cval=(0, 255),
            mode=ia.ALL
        )),

        #
        # Execute 0 to 5 of the following (less important) augmenters per
        # image. Don't execute all of them, as that would often be way too
        # strong.
        #
        iaa.SomeOf((0, 5),
            [
                # Convert some images into their superpixel representation,
                # sample between 20 and 200 superpixels per image, but do
                # not replace all superpixels with their average, only
                # some of them (p_replace).
                sometimes(
                    iaa.Superpixels(
                        p_replace=(0, 1.0),
                        n_segments=(20, 200)
                    )
                ),

                # Blur each image with varying strength using
                # gaussian blur (sigma between 0 and 3.0),
                # average/uniform blur (kernel size between 2x2 and 7x7)
                # median blur (kernel size between 3x3 and 11x11).
                iaa.OneOf([
                    iaa.GaussianBlur((0, 3.0)),
                    iaa.AverageBlur(k=(2, 7)),
                    iaa.MedianBlur(k=(3, 11)),
                ]),

                # Sharpen each image, overlay the result with the original
                # image using an alpha between 0 (no sharpening) and 1
                # (full sharpening effect).
                iaa.Sharpen(alpha=(0, 1.0), lightness=(0.75, 1.5)),

                # Same as sharpen, but for an embossing effect.
                iaa.Emboss(alpha=(0, 1.0), strength=(0, 2.0)),

                # Search in some images either for all edges or for
                # directed edges. These edges are then marked in a black
                # and white image and overlayed with the original image
                # using an alpha of 0 to 0.7.
                sometimes(iaa.OneOf([
                    iaa.EdgeDetect(alpha=(0, 0.7)),
                    iaa.DirectedEdgeDetect(
                        alpha=(0, 0.7), direction=(0.0, 1.0)
                    ),
                ])),

                # Add gaussian noise to some images.
                # In 50% of these cases, the noise is randomly sampled per
                # channel and pixel.
                # In the other 50% of all cases it is sampled once per
                # pixel (i.e. brightness change).
                iaa.AdditiveGaussianNoise(
                    loc=0, scale=(0.0, 0.05*255), per_channel=0.5
                ),

                # Either drop randomly 1 to 10% of all pixels (i.e. set
                # them to black) or drop them on an image with 2-5% percent
                # of the original size, leading to large dropped
                # rectangles.
                iaa.OneOf([
                    iaa.Dropout((0.01, 0.1), per_channel=0.5),
                    iaa.CoarseDropout(
                        (0.03, 0.15), size_percent=(0.02, 0.05),
                        per_channel=0.2
                    ),
                ]),

                # Invert each image's channel with 5% probability.
                # This sets each pixel value v to 255-v.
                iaa.Invert(0.05, per_channel=True), # invert color channels

                # Add a value of -10 to 10 to each pixel.
                iaa.Add((-10, 10), per_channel=0.5),

                # Change brightness of images (50-150% of original value).
                iaa.Multiply((0.5, 1.5), per_channel=0.5),

                # Improve or worsen the contrast of images.
                iaa.LinearContrast((0.5, 2.0), per_channel=0.5),

                # Convert each image to grayscale and then overlay the
                # result with the original with random alpha. I.e. remove
                # colors with varying strengths.
                iaa.Grayscale(alpha=(0.0, 1.0)),

                # In some images move pixels locally around (with random
                # strengths).
                sometimes(
                    iaa.ElasticTransformation(alpha=(0.5, 3.5), sigma=0.25)
                ),

                # In some images distort local areas with varying strength.
                sometimes(iaa.PiecewiseAffine(scale=(0.01, 0.05)))
            ],
             # do all of the above augmentations in random order
             random_order=True
         )
    ],
    # do all of the above augmentations in random order
    random_order=True
)


# In[13]:


#deep augmenter
@MultiClass
def DeepImgAugDataGeneratorMC(list_files, 
        id_label_map, 
        batch_size, 
        augment=False,
        shuf = False):
    
    sometimes = lambda aug: iaa.Sometimes(0.5, aug)
    return iaa.Sequential(
    [
        #
        # Apply the following augmenters to most images.
        #
        iaa.Fliplr(0.5), # horizontally flip 50% of all images
        iaa.Flipud(0.2), # vertically flip 20% of all images

        # crop some of the images by 0-10% of their height/width
        sometimes(iaa.Crop(percent=(0, 0.1))),

        # Apply affine transformations to some of the images
        # - scale to 80-120% of image height/width (each axis independently)
        # - translate by -20 to +20 relative to height/width (per axis)
        # - rotate by -45 to +45 degrees
        # - shear by -16 to +16 degrees
        # - order: use nearest neighbour or bilinear interpolation (fast)
        # - mode: use any available mode to fill newly created pixels
        #         see API or scikit-image for which modes are available
        # - cval: if the mode is constant, then use a random brightness
        #         for the newly created pixels (e.g. sometimes black,
        #         sometimes white)
        sometimes(iaa.Affine(
            scale={"x": (0.8, 1.2), "y": (0.8, 1.2)},
            translate_percent={"x": (-0.2, 0.2), "y": (-0.2, 0.2)},
            rotate=(-45, 45),
            shear=(-16, 16),
            order=[0, 1],
            cval=(0, 255),
            mode=ia.ALL
        )),

        #
        # Execute 0 to 5 of the following (less important) augmenters per
        # image. Don't execute all of them, as that would often be way too
        # strong.
        #
        iaa.SomeOf((0, 5),
            [
                # Convert some images into their superpixel representation,
                # sample between 20 and 200 superpixels per image, but do
                # not replace all superpixels with their average, only
                # some of them (p_replace).
                sometimes(
                    iaa.Superpixels(
                        p_replace=(0, 1.0),
                        n_segments=(20, 200)
                    )
                ),

                # Blur each image with varying strength using
                # gaussian blur (sigma between 0 and 3.0),
                # average/uniform blur (kernel size between 2x2 and 7x7)
                # median blur (kernel size between 3x3 and 11x11).
                iaa.OneOf([
                    iaa.GaussianBlur((0, 3.0)),
                    iaa.AverageBlur(k=(2, 7)),
                    iaa.MedianBlur(k=(3, 11)),
                ]),

                # Sharpen each image, overlay the result with the original
                # image using an alpha between 0 (no sharpening) and 1
                # (full sharpening effect).
                iaa.Sharpen(alpha=(0, 1.0), lightness=(0.75, 1.5)),

                # Same as sharpen, but for an embossing effect.
                iaa.Emboss(alpha=(0, 1.0), strength=(0, 2.0)),

                # Search in some images either for all edges or for
                # directed edges. These edges are then marked in a black
                # and white image and overlayed with the original image
                # using an alpha of 0 to 0.7.
                sometimes(iaa.OneOf([
                    iaa.EdgeDetect(alpha=(0, 0.7)),
                    iaa.DirectedEdgeDetect(
                        alpha=(0, 0.7), direction=(0.0, 1.0)
                    ),
                ])),

                # Add gaussian noise to some images.
                # In 50% of these cases, the noise is randomly sampled per
                # channel and pixel.
                # In the other 50% of all cases it is sampled once per
                # pixel (i.e. brightness change).
                iaa.AdditiveGaussianNoise(
                    loc=0, scale=(0.0, 0.05*255), per_channel=0.5
                ),

                # Either drop randomly 1 to 10% of all pixels (i.e. set
                # them to black) or drop them on an image with 2-5% percent
                # of the original size, leading to large dropped
                # rectangles.
                iaa.OneOf([
                    iaa.Dropout((0.01, 0.1), per_channel=0.5),
                    iaa.CoarseDropout(
                        (0.03, 0.15), size_percent=(0.02, 0.05),
                        per_channel=0.2
                    ),
                ]),

                # Invert each image's channel with 5% probability.
                # This sets each pixel value v to 255-v.
                iaa.Invert(0.05, per_channel=True), # invert color channels

                # Add a value of -10 to 10 to each pixel.
                iaa.Add((-10, 10), per_channel=0.5),

                # Change brightness of images (50-150% of original value).
                iaa.Multiply((0.5, 1.5), per_channel=0.5),

                # Improve or worsen the contrast of images.
                iaa.LinearContrast((0.5, 2.0), per_channel=0.5),

                # Convert each image to grayscale and then overlay the
                # result with the original with random alpha. I.e. remove
                # colors with varying strengths.
                iaa.Grayscale(alpha=(0.0, 1.0)),

                # In some images move pixels locally around (with random
                # strengths).
                sometimes(
                    iaa.ElasticTransformation(alpha=(0.5, 3.5), sigma=0.25)
                ),

                # In some images distort local areas with varying strength.
                sometimes(iaa.PiecewiseAffine(scale=(0.01, 0.05)))
            ],
             # do all of the above augmentations in random order
             random_order=True
         )
    ],
    # do all of the above augmentations in random order
    random_order=True
)


# In[14]:


def classic_train_datagen(fromfFoldTrain,
                          image_sizey_,
                          image_sizex_,
                          batch_size_,
                          shuffle_ = True,
                          ):
    
    train_generator = ImageDataGenerator(rescale=1/255.0,
                                    rotation_range=20,
                                    zoom_range=0.05,
                                    width_shift_range=0.1,
                                    height_shift_range=0.1,
                                    shear_range=0.05,
                                    horizontal_flip=True,
                                    vertical_flip=True,
                                    fill_mode="nearest",
                                    )                      
    return train_generator.flow_from_directory(
        fromfFoldTrain,
        target_size=(image_sizey_, image_sizex_),
        color_mode="rgb",
        batch_size= batch_size_,
        shuffle=shuffle_,
        class_mode='categorical')


def classic_val_datagen(fromfFoldValid,
                        image_sizey_,
                        image_sizex_,
                        batch_size_,
                        shuffle_ = False,
                        ):
    
    simple_generator = ImageDataGenerator(rescale=1/255.0)
    return simple_generator.flow_from_directory(
        fromfFoldValid,
        target_size=(image_sizey_, image_sizex_),
        color_mode="rgb",
        batch_size = batch_size_,
        shuffle=shuffle_,
        class_mode='categorical')


# In[17]:


import os
module_name = 'bvv_utils'

