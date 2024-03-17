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





### 编码器的架构

将输入评论输入到embedding中，随后将多个评论flatten到一维数组，接着，通过一个全连接层（Linear层）将输出转换成预测值。这一过程首先通过GRU网络处理经过embedding层的评论文本，利用双向GRU可以获取更加丰富的上下文信息，增强模型对文本的理解能力。双向GRU输出的结果会经过一个dropout层，以减少模型过拟合的风险。

随后，将双向GRU的输出在最后一个维度上相加，这样做是为了将来自两个方向的信息合并成一个统一的表示。接下来，这个合并后的输出被重塑成预定的形状，以适应全连接层的输入要求。全连接层的作用是将这些复杂的表示转换成最终的预测值，这些预测值可以用于分类或回归任务。

模型的最后输出是通过一个squeeze操作去除最后一个维度，因为这个维度在之前的全连接层输出中仅用于保持数据形状的一致性，并不包含实际的预测信息。最终，模型输出的是每个评论文本的预测值，这可以直接用于后续的评估或应用中。

通过这个过程，模型能够有效地处理和理解大量的文本数据，将复杂的文本信息转换成可用于决策的具体数值，为文本分析、情感分析等应用提供了强大的工具。

### 解码器的架构

解码器的架构采用了Luong注意力机制，以提高模型对输入序列的理解和输出序列的生成能力。该架构首先将输入序列通过一个嵌入层（embedding layer），并应用dropout以减少过拟合风险。接着，使用GRU网络处理嵌入后的序列，同时采用dropout进一步减少过拟合。GRU的输出既用于计算注意力权重，也用于后续的序列生成。

在注意力机制部分，解码器使用当前GRU状态和所有编码器输出来计算注意力权重，然后将这些权重应用于编码器输出以获得加权平均，形成上下文向量。这一过程分别针对两个评论的编码器输出进行，生成两个上下文向量。

随后，将GRU的输出和两个上下文向量合并，通过一个全连接层（concat）进行转换，以融合来自不同源的信息。这个合并后的输出经过另一个全连接层（out），转换成最终的预测输出，通常用于生成序列的下一个元素。

最终，解码器输出包括最终预测输出、更新后的隐藏状态、以及用户和商家的注意力权重。这些输出不仅为序列生成提供了基础，还允许我们可视化模型如何聚焦于输入序列的不同部分，以理解其决策过程。

### 模型的输入和输出

模型在训练时每一次通过解码器输入两条浓缩的用户评论，输出两个用户评论对应的标签。其中标签来源于chatgpt对于两个评论的总结。

经过训练后，只要输入任意两个该领域的评论，模型都能输出若干个对应的标签。

对于每一个商品，通过Kmeans算法将所有标签对应的词向量进行聚类，选取n个聚类中心作为该商品最后的标签



### 与Transformer相比，Ref2Label模型的优势主要体现在以下几个方面：

1. **资源效率**：Ref2Label模型由于采用了较为轻量的RNN结构，其在参数数量和计算资源消耗上都相对较低，使得模型更适合在资源受限的环境下运行。
2. **响应速度**：得益于其简洁的架构，Ref2Label模型在处理每一条评论时的响应时间更短，特别适合需要快速处理大量数据的场景。
3. **可解释性**：通过引入Luong注意力机制，Ref2Label模型在聚焦于文本关键信息的同时，也提供了一定程度的可解释性，有助于理解模型是如何从整体评论中提取出关键信息来生成标签的。
4. **特定任务的定制化**：Ref2Label模型针对的是将长评论浓缩成简短标签的特定任务，其设计和优化都是为了更好地完成这一任务。相比之下，Transformer模型虽然强大且通用，但在处理这类特定且相对简单的任务时，可能存在过度设计的问题。

#### 与金融结合

我的毕业设计搭建了一个电商系统，用户可以购买商品，系统也可以推荐商品。但是毕业设计要求解决计算机科学与金融学交叉的复杂问题。所以我要进一步改造电商系统，使其与金融学交叉。

如果我要引入一个fine-tuning的大语言模型或者构造一个电商金融领域的知识库解决用户金融方面的问题。我可以做哪些模块？给出一些选项和具体一些可行的实施方案。
