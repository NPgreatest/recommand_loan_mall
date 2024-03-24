import PyPDF2

def count_chars_in_pdf(pdf_path):
    char_count = 0
    try:
        with open(pdf_path, 'rb') as file:
            reader = PyPDF2.PdfReader(file)
            for page in reader.pages:
                text = page.extract_text()
                if text:
                    char_count += len(text)
    except Exception as e:
        print(f"Error reading file: {e}")
        return None
    return char_count

# 替换为你的 PDF 文件路径
pdf_path = 'emnlp19a - count.pdf'
print(f"Total number of characters in the PDF: {count_chars_in_pdf(pdf_path)}")

