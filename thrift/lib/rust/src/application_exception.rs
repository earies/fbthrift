/*
 * Copyright 2019-present Facebook, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

use crate::deserialize::Deserialize;
use crate::protocol::{ProtocolReader, ProtocolWriter};
use crate::serialize::Serialize;
use crate::thrift_protocol::ProtocolID;
use crate::ttype::TType;
use crate::Result;

#[derive(Debug, Copy, Clone, Eq, PartialEq, Ord, PartialOrd, Hash)]
#[repr(i32)]
pub enum ApplicationExceptionErrorCode {
    Unknown = 0,
    UnknownMethod = 1,
    InvalidMessageType = 2,
    WrongMethodName = 3,
    BadSequenceID = 4,
    MissingResult = 5,
    InternalError = 6,
    ProtocolError = 7,
    InvalidTransform = 8,
    InvalidProtocol = 9,
    UnsupportedClientType = 10,
    Loadshedding = 11,
    Timeout = 12,
    InjectedFailure = 13,
}

impl Default for ApplicationExceptionErrorCode {
    fn default() -> Self {
        ApplicationExceptionErrorCode::Unknown
    }
}

const TAPPLICATION_EXCEPTION_ERROR_CODE: &'static str = "ApplicationExceptionErrorCode";

/// ApplicationException is *not* actually an error type in the Rust sense, but
/// a Thrift-specific serializable
#[derive(Debug, Clone, Eq, PartialEq, Ord, PartialOrd, Hash, Default)]
pub struct ApplicationException {
    pub message: String,
    pub type_: ApplicationExceptionErrorCode,
}

impl ApplicationException {
    #[cold]
    pub fn new(type_: ApplicationExceptionErrorCode, message: String) -> Self {
        ApplicationException { message, type_ }
    }

    // Used for methods which should exist but don't
    #[cold]
    pub fn unimplemented_method(handler: &str, method: &str) -> Self {
        ApplicationException {
            type_: ApplicationExceptionErrorCode::UnknownMethod,
            message: format!("Method {} not implemented for {}", method, handler),
        }
    }

    /// Indicator that this service doesn't have the method, but another might
    #[inline]
    pub fn unknown_method() -> Self {
        ApplicationException {
            type_: ApplicationExceptionErrorCode::UnknownMethod,
            message: String::new(), // no allocation
        }
    }

    #[cold]
    pub fn missing_arg(method: &str, arg: &str) -> Self {
        ApplicationException {
            type_: ApplicationExceptionErrorCode::ProtocolError,
            message: format!("{} missing arg {}", method, arg),
        }
    }

    #[cold]
    pub fn missing_field(method: &str, field: &str) -> Self {
        ApplicationException {
            type_: ApplicationExceptionErrorCode::ProtocolError,
            message: format!("Struct {} missing field {}", method, field),
        }
    }

    #[cold]
    pub fn invalid_protocol(badproto: ProtocolID) -> Self {
        ApplicationException {
            type_: ApplicationExceptionErrorCode::InvalidProtocol,
            message: format!("Invalid protocol {:?}", badproto),
        }
    }
}

impl<P> Deserialize<P> for ApplicationException
where
    P: ProtocolReader,
{
    /// Decodes a ApplicationException message from a Protocol stream
    fn read(iprot: &mut P) -> Result<Self> {
        iprot.read_struct_begin(|_| ())?;

        let mut message = String::from("");
        let mut type_ = ApplicationExceptionErrorCode::Unknown;

        loop {
            // Start reading the next field
            let (_, ttype, id) = iprot.read_field_begin(|_| ())?;

            match (ttype, id) {
                (TType::Stop, _) => break,
                (TType::String, 1) => message = iprot.read_string()?,
                (TType::I32, 2) => {
                    type_ = match iprot.read_i32()? {
                        1 => ApplicationExceptionErrorCode::UnknownMethod,
                        2 => ApplicationExceptionErrorCode::InvalidMessageType,
                        3 => ApplicationExceptionErrorCode::WrongMethodName,
                        4 => ApplicationExceptionErrorCode::BadSequenceID,
                        5 => ApplicationExceptionErrorCode::MissingResult,
                        6 => ApplicationExceptionErrorCode::InternalError,
                        7 => ApplicationExceptionErrorCode::ProtocolError,
                        8 => ApplicationExceptionErrorCode::InvalidTransform,
                        9 => ApplicationExceptionErrorCode::InvalidProtocol,
                        10 => ApplicationExceptionErrorCode::UnsupportedClientType,
                        11 => ApplicationExceptionErrorCode::Loadshedding,
                        12 => ApplicationExceptionErrorCode::Timeout,
                        13 => ApplicationExceptionErrorCode::InjectedFailure,

                        _ => ApplicationExceptionErrorCode::Unknown,
                    }
                }
                (ttype, _) => iprot.skip(ttype)?,
            };

            // Finished reading the end of the field
            iprot.read_field_end()?;
        }
        iprot.read_struct_end()?;
        return Ok(ApplicationException::new(type_, message));
    }
}

impl<'a, P> Serialize<P> for &'a ApplicationException
where
    P: ProtocolWriter,
{
    /// Writes an application exception to the Protocol stream
    fn write(self, oprot: &mut P) {
        oprot.write_struct_begin(TAPPLICATION_EXCEPTION_ERROR_CODE);

        if self.message.len() > 0 {
            oprot.write_field_begin("message", TType::String, 1);
            oprot.write_string(&self.message);
            oprot.write_field_end();
        }
        oprot.write_field_begin("type", TType::I32, 2);
        oprot.write_i32(self.type_ as i32);
        oprot.write_field_end();
        oprot.write_field_stop();
        oprot.write_struct_end();
    }
}
