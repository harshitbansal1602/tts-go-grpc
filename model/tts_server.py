import grpc 
import tts_pb2
import tts_pb2_grpc
from concurrent import futures
import tts_bark
import google.protobuf.empty_pb2

class BridgeServicer(tts_pb2_grpc.BridgeServicer):
    def __init__(self) -> None:
        super().__init__()
    
    def GetSpeechStream(self, request:tts_pb2.Text, context):
        return
    
    def GetSpeech(self, request: tts_pb2.Text, context):
        model = tts_bark.BarkTTS()
        speechArray = model.multiline_tts(request.text)
        return tts_pb2.Speech(speech=speechArray, length=speechArray.__len__())
    
    def GetPartTextToSpeech(self, request, context):
        return
    
    def DownloadBarkModel(self, request, context):
        tts_bark.BarkTTS()
        return google.protobuf.empty_pb2.Empty()

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    tts_pb2_grpc.add_BridgeServicer_to_server(BridgeServicer(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    print("Started server...")
    server.wait_for_termination()

if __name__ == "__main__":
    serve()