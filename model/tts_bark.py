import io
import numpy as np
from scipy.io.wavfile import write as write_wav

from bark.generation import (
    generate_text_semantic,
    preload_models,
)
from bark.api import semantic_to_waveform
from bark import generate_audio, SAMPLE_RATE
from processing import Processing
import time

class BarkTTS():
    def __init__(self) -> None:
        preload_models()
        self.GEN_TEMP = 0.6
        self.SPEAKER = "v2/en_speaker_6"
        self.SILENCE_TIME_SECS = 0.25
        pass

    def multiline_tts(self, text):
        processing = Processing()
        script = processing.flatten(text)
        sentences = processing.tokenize(script)

        silence = np.zeros(int(self.SILENCE_TIME_SECS * SAMPLE_RATE))  # quarter second of silence
        pieces = []
        for sentence in sentences:
            semantic_tokens = generate_text_semantic(
                sentence,
                history_prompt=self.SPEAKER,
                temp=self.GEN_TEMP,
                min_eos_p=0.05,  # this controls how likely the generation is to end
            )

            audio_array = semantic_to_waveform(semantic_tokens, history_prompt=self.SPEAKER,)
            pieces += [audio_array, silence.copy()]
        
        pieces = np.concatenate(pieces)

        wav_file = io.BytesIO()
        start = time.time()
        print("Started making wav file")
        write_wav(wav_file, SAMPLE_RATE, pieces)
        print(time.time() - start)
        wav_file.seek(0)
        return wav_file.read()
    
    def write_to_disk(self, audio_array, filename):
        write_wav("../outputs/" + str(filename) + ".wav", SAMPLE_RATE, audio_array)
        return



