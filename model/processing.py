import nltk

class Processing():
    def __init__(self) -> None:
        pass

    def flatten(self, text, delimeter):
        return text.replace("\n", " ").strip()
    
    def tokenize(self, text):
        return nltk.sent_tokenize(text)
    

