if __name__ == '__main__':
    book = input()
    want = input()

    book_word = ""

    for i in book:
        if i.isalpha():
            book_word = book_word + i
    
    if want in book_word:
        print(1)
    else:
        print(0)