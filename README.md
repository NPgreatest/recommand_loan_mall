# Graduation Design Project Introduction

## Project Overview

> This is a mall project with front-end and back-end separated architecture, the back-end is implemented using **Golang**, and the front-end adopts the popular **Vue.js** framework. The data processing part mainly relies on **Python** and **Jupyter Notebook**.
>
> The database part uses **Mysql** to store structured data and **Postgresql** to store vectorized data using pgvector.

这是一个具有前后端分离架构的商城项目，后端使用 **Golang** 实现，前端则采用了流行的 **Vue.js** 框架。数据处理部分主要依赖于 **Python** 和 **Jupyter Notebook**。

数据库部分采用**Mysql**存储结构化数据，**Postgresql**利用pgvector存储向量化数据。

### 数据来源

项目的数据来源于UCSD公开的亚马逊数据源。您可以通过以下链接访问数据集：

[UCSD Amazon Data Source](https://cseweb.ucsd.edu/~jmcauley/datasets/amazon_v2/)

引用来源：*Empirical Methods in Natural Language Processing (EMNLP), 2019*

### 数据处理

对于原始数据，我们进行了清洗和整理，将商品信息划分为三级标签分类。此外，为了更好的展示效果，我们使用了 **DeepL API** 将这些标签翻译成中文。

## 技术实现

### 词嵌入与数据库存储

我们对海量数据进行了词向量嵌入处理，并将结果存储到 **PostgreSQL-Pgvector** 数据库中。

### Kmeans聚类

项目中对20个一级分类进行了Kmeans聚类测试。以下是聚类测试的结果展示：

<img src="/data-process/kmeans.png" alt="Kmeans聚类结果" style="zoom:50%;" />

![聚类数据可视化](/data-process/cluster.png)

## 系统亮点

- **大语言模型应用**：系统通过大语言模型接收用户的输入。
- **Fine-tuning模型**：结合fine-tuning模型，将用户输入拆分成所需的商品信息。
- **词向量嵌入**：通过词向量嵌入技术搜索出符合商品信息的对应商品。
- **Ref2Label模型**：利用Pytorch训练的Ref2Label的大语言模型，该模型基于Transformer架构，结合了Kmeans聚类算法，构建了一个能够持续自我更新的智能推荐系统。通过对大量用户评论的深入分析和处理，ref2label模型能够自动提炼出能够最准确反映商品特性的标签，进而利用这些标签为用户推荐最合适的商品。这一过程不仅显著提升了推荐的准确性，也极大地丰富了用户的购物体验。模型通过细致的fine-tuning和词嵌入技术的应用，进一步增强了推荐系统的性能。

---

基于fine-tuning, embedding, ML的初步构想流程图
![流程图](/data-process/naive-target.png)

该毕业设计项目不仅体现了现代前端和后端技术的结合，而且展示了数据处理和机器学习在电商领域的应用。这不仅是一次技术挑战，也是对创新应用的探索。



### 实现的效果

<img src="/pic/ai1.png" alt="流程图" style="zoom: 50%;" /><img src="/pic/ai2.png" alt="流程图" style="zoom: 50%;" /><img src="/pic/review.png" alt="流程图" style="zoom: 50%;" />



总体推荐系统架构

<img src="/pic/ref2label.drawio.png" alt="流程图" style="zoom: 50%;" />





数据库E-R图

<img src="/pic/e-r.png" alt="流程图" style="zoom: 50%;" />



* 标签聚类结果

<img src="/pic/word-cluster.png" alt="流程图" style="zoom: 150%;" />

<img src="/pic/words.png" alt="流程图" style="zoom: 150%;" />

* 2024/1/24: 后端模块新增用户金融数据存取模块，用户可以输入自己的财务信息，以供之后的预测；利用chatgpt的api接口对每个商品进行精准定价，用于后面的机器学习。
* 2024/1/25: 初步完成fine-tuning推荐系统以及分组embedding获取候选商品的过程。调整了所有商品的定价。下一步进行机器学习。
* 2024/1/26:清洗并且导入了所有商品评论的数据，并且将查看评论功能添加到后端。除此以外，添加上传头像的功能。
* 2024/2/10:完成深度学习ref2label模块，撰写开题报告