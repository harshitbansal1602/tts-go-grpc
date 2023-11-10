from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Text(_message.Message):
    __slots__ = ["text", "part"]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    PART_FIELD_NUMBER: _ClassVar[int]
    text: str
    part: int
    def __init__(self, text: _Optional[str] = ..., part: _Optional[int] = ...) -> None: ...

class Speech(_message.Message):
    __slots__ = ["speech", "length"]
    SPEECH_FIELD_NUMBER: _ClassVar[int]
    LENGTH_FIELD_NUMBER: _ClassVar[int]
    speech: bytes
    length: int
    def __init__(self, speech: _Optional[bytes] = ..., length: _Optional[int] = ...) -> None: ...
