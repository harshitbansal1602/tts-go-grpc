<!-- python grpc init script  -->
python -m grpc_tools.protoc -I=<PROTO_FIlE_PATH> --python_out=<MODEL_PATH> --pyi_out=<MODEL_PATH> --grpc_python_out=<MODEL_PATH> <PROTO_FIlE_PATH>

