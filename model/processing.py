import nltk

class Processing():
    def __init__(self) -> None:
        pass

    def flatten(self, text, delimeter="\n"):
        return text.replace(delimeter, " ").strip()
    
    def tokenize(self, text):
        return nltk.sent_tokenize(text)
    

