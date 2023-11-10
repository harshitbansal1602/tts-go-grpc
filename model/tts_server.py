import grpc 
import tts_pb2
import tts_pb2_grpc
from concurrent import futures

class BridgeServicer(tts_pb2_grpc.BridgeServicer):
    def __init__(self) -> None:
        super().__init__()
    
    def GetSpeechStream(self, request:tts_pb2.Text, context):
        return
    
    def GetSpeech(self, request: tts_pb2.Text, context):
        return
    
    def GetPartTextToSpeech(self, request, context):
        return
    
    def DownloadModels(self, request, context):
        return

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    tts_pb2_grpc.add_BridgeServicer_to_server(BridgeServicer(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    serve()