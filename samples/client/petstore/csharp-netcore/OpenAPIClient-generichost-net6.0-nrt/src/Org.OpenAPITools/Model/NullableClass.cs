/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Org.OpenAPITools.Client.OpenAPIDateConverter;
using OpenAPIClientUtils = Org.OpenAPITools.Client.ClientUtils;

namespace Org.OpenAPITools.Model
{
    /// <summary>
    /// NullableClass
    /// </summary>
    [DataContract(Name = "NullableClass")]
    public partial class NullableClass : Dictionary<String, Object>, IEquatable<NullableClass>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="NullableClass" /> class.
        /// </summary>
        /// <param name="integerProp">integerProp.</param>
        /// <param name="numberProp">numberProp.</param>
        /// <param name="booleanProp">booleanProp.</param>
        /// <param name="stringProp">stringProp.</param>
        /// <param name="dateProp">dateProp.</param>
        /// <param name="datetimeProp">datetimeProp.</param>
        /// <param name="arrayNullableProp">arrayNullableProp.</param>
        /// <param name="arrayAndItemsNullableProp">arrayAndItemsNullableProp.</param>
        /// <param name="arrayItemsNullable">arrayItemsNullable.</param>
        /// <param name="objectNullableProp">objectNullableProp.</param>
        /// <param name="objectAndItemsNullableProp">objectAndItemsNullableProp.</param>
        /// <param name="objectItemsNullable">objectItemsNullable.</param>
        public NullableClass(int? integerProp = default, decimal? numberProp = default, bool? booleanProp = default, string? stringProp = default, DateTime? dateProp = default, DateTime? datetimeProp = default, List<Object>? arrayNullableProp = default, List<Object>? arrayAndItemsNullableProp = default, List<Object>? arrayItemsNullable = default, Dictionary<string, Object>? objectNullableProp = default, Dictionary<string, Object>? objectAndItemsNullableProp = default, Dictionary<string, Object>? objectItemsNullable = default) : base()
        {
            this.IntegerProp = integerProp;
            this.NumberProp = numberProp;
            this.BooleanProp = booleanProp;
            this.StringProp = stringProp;
            this.DateProp = dateProp;
            this.DatetimeProp = datetimeProp;
            this.ArrayNullableProp = arrayNullableProp;
            this.ArrayAndItemsNullableProp = arrayAndItemsNullableProp;
            this.ArrayItemsNullable = arrayItemsNullable;
            this.ObjectNullableProp = objectNullableProp;
            this.ObjectAndItemsNullableProp = objectAndItemsNullableProp;
            this.ObjectItemsNullable = objectItemsNullable;
        }

        /// <summary>
        /// Gets or Sets IntegerProp
        /// </summary>
        [DataMember(Name = "integer_prop", EmitDefaultValue = true)]
        public int? IntegerProp { get; set; }

        /// <summary>
        /// Gets or Sets NumberProp
        /// </summary>
        [DataMember(Name = "number_prop", EmitDefaultValue = true)]
        public decimal? NumberProp { get; set; }

        /// <summary>
        /// Gets or Sets BooleanProp
        /// </summary>
        [DataMember(Name = "boolean_prop", EmitDefaultValue = true)]
        public bool? BooleanProp { get; set; }

        /// <summary>
        /// Gets or Sets StringProp
        /// </summary>
        [DataMember(Name = "string_prop", EmitDefaultValue = true)]
        public string? StringProp { get; set; }

        /// <summary>
        /// Gets or Sets DateProp
        /// </summary>
        [DataMember(Name = "date_prop", EmitDefaultValue = true)]
        [JsonConverter(typeof(OpenAPIDateConverter))]
        public DateTime? DateProp { get; set; }

        /// <summary>
        /// Gets or Sets DatetimeProp
        /// </summary>
        [DataMember(Name = "datetime_prop", EmitDefaultValue = true)]
        public DateTime? DatetimeProp { get; set; }

        /// <summary>
        /// Gets or Sets ArrayNullableProp
        /// </summary>
        [DataMember(Name = "array_nullable_prop", EmitDefaultValue = true)]
        public List<Object>? ArrayNullableProp { get; set; }

        /// <summary>
        /// Gets or Sets ArrayAndItemsNullableProp
        /// </summary>
        [DataMember(Name = "array_and_items_nullable_prop", EmitDefaultValue = true)]
        public List<Object>? ArrayAndItemsNullableProp { get; set; }

        /// <summary>
        /// Gets or Sets ArrayItemsNullable
        /// </summary>
        [DataMember(Name = "array_items_nullable", EmitDefaultValue = false)]
        public List<Object>? ArrayItemsNullable { get; set; }

        /// <summary>
        /// Gets or Sets ObjectNullableProp
        /// </summary>
        [DataMember(Name = "object_nullable_prop", EmitDefaultValue = true)]
        public Dictionary<string, Object>? ObjectNullableProp { get; set; }

        /// <summary>
        /// Gets or Sets ObjectAndItemsNullableProp
        /// </summary>
        [DataMember(Name = "object_and_items_nullable_prop", EmitDefaultValue = true)]
        public Dictionary<string, Object>? ObjectAndItemsNullableProp { get; set; }

        /// <summary>
        /// Gets or Sets ObjectItemsNullable
        /// </summary>
        [DataMember(Name = "object_items_nullable", EmitDefaultValue = false)]
        public Dictionary<string, Object>? ObjectItemsNullable { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class NullableClass {\n");
            sb.Append("  ").Append(base.ToString().Replace("\n", "\n  ")).Append("\n");
            sb.Append("  IntegerProp: ").Append(IntegerProp).Append("\n");
            sb.Append("  NumberProp: ").Append(NumberProp).Append("\n");
            sb.Append("  BooleanProp: ").Append(BooleanProp).Append("\n");
            sb.Append("  StringProp: ").Append(StringProp).Append("\n");
            sb.Append("  DateProp: ").Append(DateProp).Append("\n");
            sb.Append("  DatetimeProp: ").Append(DatetimeProp).Append("\n");
            sb.Append("  ArrayNullableProp: ").Append(ArrayNullableProp).Append("\n");
            sb.Append("  ArrayAndItemsNullableProp: ").Append(ArrayAndItemsNullableProp).Append("\n");
            sb.Append("  ArrayItemsNullable: ").Append(ArrayItemsNullable).Append("\n");
            sb.Append("  ObjectNullableProp: ").Append(ObjectNullableProp).Append("\n");
            sb.Append("  ObjectAndItemsNullableProp: ").Append(ObjectAndItemsNullableProp).Append("\n");
            sb.Append("  ObjectItemsNullable: ").Append(ObjectItemsNullable).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return OpenAPIClientUtils.compareLogic.Compare(this, input as NullableClass).AreEqual;
        }

        /// <summary>
        /// Returns true if NullableClass instances are equal
        /// </summary>
        /// <param name="input">Instance of NullableClass to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(NullableClass input)
        {
            return OpenAPIClientUtils.compareLogic.Compare(this, input).AreEqual;
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = base.GetHashCode();
                if (this.IntegerProp != null)
                {
                    hashCode = (hashCode * 59) + this.IntegerProp.GetHashCode();
                }
                if (this.NumberProp != null)
                {
                    hashCode = (hashCode * 59) + this.NumberProp.GetHashCode();
                }
                if (this.BooleanProp != null)
                {
                    hashCode = (hashCode * 59) + this.BooleanProp.GetHashCode();
                }
                if (this.StringProp != null)
                {
                    hashCode = (hashCode * 59) + this.StringProp.GetHashCode();
                }
                if (this.DateProp != null)
                {
                    hashCode = (hashCode * 59) + this.DateProp.GetHashCode();
                }
                if (this.DatetimeProp != null)
                {
                    hashCode = (hashCode * 59) + this.DatetimeProp.GetHashCode();
                }
                if (this.ArrayNullableProp != null)
                {
                    hashCode = (hashCode * 59) + this.ArrayNullableProp.GetHashCode();
                }
                if (this.ArrayAndItemsNullableProp != null)
                {
                    hashCode = (hashCode * 59) + this.ArrayAndItemsNullableProp.GetHashCode();
                }
                if (this.ArrayItemsNullable != null)
                {
                    hashCode = (hashCode * 59) + this.ArrayItemsNullable.GetHashCode();
                }
                if (this.ObjectNullableProp != null)
                {
                    hashCode = (hashCode * 59) + this.ObjectNullableProp.GetHashCode();
                }
                if (this.ObjectAndItemsNullableProp != null)
                {
                    hashCode = (hashCode * 59) + this.ObjectAndItemsNullableProp.GetHashCode();
                }
                if (this.ObjectItemsNullable != null)
                {
                    hashCode = (hashCode * 59) + this.ObjectItemsNullable.GetHashCode();
                }
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        public IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
