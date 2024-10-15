/*
 * Copyright 2024 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

namespace go openapi

struct _ServiceOptions {
      1:required Document document
}

struct _StructOptions {
      1:required Schema schema
}

struct _MethodOptions {
      1:required Operation operation
}

struct _FieldOptions {
      1:required Parameter parameter
      2:required Schema property
}

struct AdditionalPropertiesItem {
  1: SchemaOrReference schema_or_reference,
  2: bool boolean
}

struct Any {
  1: _Any value,
  2: string yaml
}

struct _Any {
  1: string type_url,
  2: binary value
}

struct AnyOrExpression {
  1: Any any,
  2: Expression expression
}

struct Callback {
  1: list<NamedPathItem> path,
  2: list<NamedAny> specification_extension
}

struct CallbackOrReference {
  1: Callback callback,
  2: Reference reference
}

struct CallbacksOrReferences {
  1: list<NamedCallbackOrReference> additional_properties
}

struct Components {
  1: SchemasOrReferences schemas,
  2: ResponsesOrReferences responses,
  3: ParametersOrReferences parameters,
  4: ExamplesOrReferences examples,
  5: RequestBodiesOrReferences request_bodies,
  6: HeadersOrReferences headers,
  7: SecuritySchemesOrReferences security_schemes,
  8: LinksOrReferences links,
  9: CallbacksOrReferences callbacks,
  10: list<NamedAny> specification_extension
}

struct Contact {
  1: string name,
  2: string url,
  3: string email,
  4: list<NamedAny> specification_extension
}

struct DefaultType {
  1: double number,
  2: bool boolean,
  3: string string
}

struct Discriminator {
  1: string property_name,
  2: Strings mapping,
  3: list<NamedAny> specification_extension
}

struct Document {
  1: string openapi,
  2: Info info,
  3: list<Server> servers,
  4: Paths paths,
  5: Components components,
  6: list<SecurityRequirement> security,
  7: list<Tag> tags,
  8: ExternalDocs external_docs,
  9: list<NamedAny> specification_extension
}

struct Encoding {
  1: string content_type,
  2: HeadersOrReferences headers,
  3: string style,
  4: bool explode,
  5: bool allow_reserved,
  6: list<NamedAny> specification_extension
}

struct Encodings {
  1: list<NamedEncoding> additional_properties
}

struct Example {
  1: string summary,
  2: string description,
  3: Any value,
  4: string external_value,
  5: list<NamedAny> specification_extension
}

struct ExampleOrReference {
  1: Example example,
  2: Reference reference
}

struct ExamplesOrReferences {
  1: list<NamedExampleOrReference> additional_properties
}

struct Expression {
  1: list<NamedAny> additional_properties
}

struct ExternalDocs {
  1: string description,
  2: string url,
  3: list<NamedAny> specification_extension
}

struct Header {
  1: string description,
  2: bool required,
  3: bool deprecated,
  4: bool allow_empty_value,
  5: string style,
  6: bool explode,
  7: bool allow_reserved,
  8: SchemaOrReference schema,
  9: Any example,
  10: ExamplesOrReferences examples,
  11: MediaTypes content,
  12: list<NamedAny> specification_extension
}

struct HeaderOrReference {
  1: Header header,
  2: Reference reference
}

struct HeadersOrReferences {
  1: list<NamedHeaderOrReference> additional_properties
}

struct Info {
  1: string title,
  2: string description,
  3: string terms_of_service,
  4: Contact contact,
  5: License license,
  6: string version,
  7: list<NamedAny> specification_extension,
  8: string summary
}

struct ItemsItem {
  1: list<SchemaOrReference> schema_or_reference
}

struct License {
  1: string name,
  2: string url,
  3: list<NamedAny> specification_extension
}

struct Link {
  1: string operation_ref,
  2: string operation_id,
  3: AnyOrExpression parameters,
  4: AnyOrExpression request_body,
  5: string description,
  6: Server server,
  7: list<NamedAny> specification_extension
}

struct LinkOrReference {
  1: Link link,
  2: Reference reference
}

struct LinksOrReferences {
  1: list<NamedLinkOrReference> additional_properties
}

struct MediaType {
  1: SchemaOrReference schema,
  2: Any example,
  3: ExamplesOrReferences examples,
  4: Encodings encoding,
  5: list<NamedAny> specification_extension
}

struct MediaTypes {
  1: list<NamedMediaType> additional_properties
}

struct NamedAny {
  1: string name,
  2: Any value
}

struct NamedCallbackOrReference {
  1: string name,
  2: CallbackOrReference value
}

struct NamedEncoding {
  1: string name,
  2: Encoding value
}

struct NamedExampleOrReference {
  1: string name,
  2: ExampleOrReference value
}

struct NamedHeaderOrReference {
  1: string name,
  2: HeaderOrReference value
}

struct NamedLinkOrReference {
  1: string name,
  2: LinkOrReference value
}

struct NamedMediaType {
  1: string name,
  2: MediaType value
}

struct NamedParameterOrReference {
  1: string name,
  2: ParameterOrReference value
}

struct NamedPathItem {
  1: string name,
  2: PathItem value
}

struct NamedRequestBodyOrReference {
  1: string name,
  2: RequestBodyOrReference value
}

struct NamedResponseOrReference {
  1: string name,
  2: ResponseOrReference value
}

struct NamedSchemaOrReference {
  1: string name,
  2: SchemaOrReference value
}

struct NamedSecuritySchemeOrReference {
  1: string name,
  2: SecuritySchemeOrReference value
}

struct NamedServerVariable {
  1: string name,
  2: ServerVariable value
}

struct NamedString {
  1: string name,
  2: string value
}

struct NamedStringArray {
  1: string name,
  2: StringArray value
}

struct OauthFlow {
  1: string authorization_url,
  2: string token_url,
  3: string refresh_url,
  4: Strings scopes,
  5: list<NamedAny> specification_extension
}

struct OauthFlows {
  1: OauthFlow implicit,
  2: OauthFlow password,
  3: OauthFlow client_credentials,
  4: OauthFlow authorization_code,
  5: list<NamedAny> specification_extension
}

struct Object {
  1: list<NamedAny> additional_properties
}

struct Operation {
  1: list<string> tags,
  2: string summary,
  3: string description,
  4: ExternalDocs external_docs,
  5: string operation_id,
  6: list<ParameterOrReference> parameters,
  7: RequestBodyOrReference request_body,
  8: Responses responses,
  9: CallbacksOrReferences callbacks,
  10: bool deprecated,
  11: list<SecurityRequirement> security,
  12: list<Server> servers,
  13: list<NamedAny> specification_extension
}

struct Parameter {
  1: string name,
  2: string in,
  3: string description,
  4: bool required,
  5: bool deprecated,
  6: bool allow_empty_value,
  7: string style,
  8: bool explode,
  9: bool allow_reserved,
  10: SchemaOrReference schema,
  11: Any example,
  12: ExamplesOrReferences examples,
  13: MediaTypes content,
  14: list<NamedAny> specification_extension
}

struct ParameterOrReference {
  1: Parameter parameter,
  2: Reference reference
}

struct ParametersOrReferences {
  1: list<NamedParameterOrReference> additional_properties
}

struct PathItem {
  1: string xref,
  2: string summary,
  3: string description,
  4: Operation get,
  5: Operation put,
  6: Operation post,
  7: Operation delete,
  8: Operation options,
  9: Operation head,
  10: Operation patch,
  11: Operation trace,
  12: list<Server> servers,
  13: list<ParameterOrReference> parameters,
  14: list<NamedAny> specification_extension
}

struct Paths {
  1: list<NamedPathItem> path
  2: list<NamedAny> specification_extension
}

struct Properties {
  1: list<NamedSchemaOrReference> additional_properties
}

struct Reference {
  1: string xref
  2: string summary
  3: string description
}

struct RequestBody {
  1: string description,
  2: MediaTypes content,
  3: bool required,
  4: list<NamedAny> specification_extension
}

struct RequestBodyOrReference {
  1: RequestBody request_body,
  2: Reference reference
}

struct RequestBodiesOrReferences {
  1: list<NamedRequestBodyOrReference> additional_properties
}

struct Response {
  1: string description,
  2: HeadersOrReferences headers,
  3: MediaTypes content,
  4: LinksOrReferences links,
  5: list<NamedAny> specification_extension
}

struct ResponseOrReference {
  1: Response response,
  2: Reference reference
}

struct Responses {
  1: ResponseOrReference default,
  2: list<NamedResponseOrReference> response_or_reference,
  3: list<NamedAny> specification_extension
}

struct ResponsesOrReferences {
  1: list<NamedResponseOrReference> additional_properties
}

struct Schema {
  1: bool nullable,
  2: Discriminator discriminator,
  3: bool read_only,
  4: bool write_only,
  5: Xml xml,
  6: ExternalDocs external_docs,
  7: Any example,
  8: bool deprecated,
  9: string title,
  10: double multiple_of,
  11: double maximum,
  12: bool exclusive_maximum,
  13: double minimum,
  14: bool exclusive_minimum,
  15: i64 max_length,
  16: i64 min_length,
  17: string pattern,
  18: i64 max_items,
  19: i64 min_items,
  20: bool unique_items,
  21: i64 max_properties,
  22: i64 min_properties,
  23: list<string> required,
  24: list<Any> enum,
  25: string type,
  26: list<SchemaOrReference> all_of,
  27: list<SchemaOrReference> one_of,
  28: list<SchemaOrReference> any_of,
  29: Schema not,
  30: ItemsItem items,
  31: Properties properties,
  32: AdditionalPropertiesItem additional_properties,
  33: DefaultType default,
  34: string description,
  35: string format,
  36: list<NamedAny> specification_extension
}

struct SchemaOrReference {
  1: Schema schema,
  2: Reference reference
}

struct SchemasOrReferences {
  1: list<NamedSchemaOrReference> additional_properties
}

struct SecurityRequirement {
  1: list<NamedStringArray> additional_properties
}

struct SecurityScheme {
  1: string _type,
  2: string description,
  3: string name,
  4: string _in,
  5: string scheme,
  6: string bearer_format,
  7: OauthFlows flows,
  8: string open_id_connect_url,
  9: list<NamedAny> specification_extension
}

struct SecuritySchemeOrReference {
  1: SecurityScheme security_scheme,
  2: Reference reference
}

struct SecuritySchemesOrReferences {
  1: list<NamedSecuritySchemeOrReference> additional_properties
}

struct Server {
  1: string url,
  2: string description,
  3: ServerVariables variables,
  4: list<NamedAny> specification_extension
}

struct ServerVariable {
  1: string _default,
  2: list<string> enum,
  3: string description,
  4: list<NamedAny> specification_extension
}

struct ServerVariables {
  1: list<NamedServerVariable> additional_properties
}

struct SpecificationExtension {
  1: double number,
  2: bool boolean,
  3: string string
}

struct StringArray {
  1: list<string> values
}

struct Strings {
  1: list<NamedString> additional_properties
}

struct Tag {
  1: string name,
  2: string description,
  3: ExternalDocs external_docs,
  4: list<NamedAny> specification_extension
}

struct Xml {
  1: string name,
  2: string namespace,
  3: string prefix,
  4: bool attribute,
  5: bool wrapped,
  6: list<NamedAny> specification_extension
}